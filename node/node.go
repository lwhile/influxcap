package node

import (
	"time"

	"strconv"

	"context"

	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/rafthttp"
	"github.com/lwhile/influxcap/backend"
)

const (
	// RUNNING status
	RUNNING Status = "running"

	// DOWNING status
	DOWNING Status = "downing"
)

const (
	heartBeatDuration = 100 * time.Millisecond
)

// Raft interface
type Raft interface {
}

// Status type
type Status string

// Conf is config of a Node
type Conf struct {
	ID    int
	Peers []string
	Join  bool
}

// Node is a influxcap instance,
// which be implemented as a raft node(Raft interface)
type Node struct {
	ID     int      `json:"id"`
	Status Status   `json:"status"`
	Join   bool     `json:"join"`
	Peers  []string `json:"peers"`

	proposeC <-chan string
	commitC  chan<- *string
	errorC   chan<- error

	stopC   chan struct{}
	Backend *backend.Influxdb

	lastIndex    uint64
	appliedIndex uint64

	// Field below is relating to raft library

	// etcd raft node interface for the commit/error channel
	node raft.Node

	// Persist storage (just memory storage now)
	storage *raft.MemoryStorage

	// Net transport
	transport *rafthttp.Transport
}

// Start a raft node
func (n *Node) Start() {
	// prepare raft config
	peers := make([]raft.Peer, len(n.Peers))
	for i := range peers {
		// The ID must not be zero
		peers[i] = raft.Peer{ID: uint64(i + 1)}
	}
	raftConf := raft.Config{
		ID:              uint64(n.ID),
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         n.storage,
		MaxSizePerMsg:   1024 * 1024,
		MaxInflightMsgs: 256,
	}
	// Invoke raft library to start a raft node
	n.node = raft.StartNode(&raftConf, peers)

	// TODO: About node status
	ss := &stats.ServerStats{}
	ss.Initialize()

	// TODO: About net transport
	n.transport = &rafthttp.Transport{
		ID:          types.ID(n.ID),
		ClusterID:   0x1000,
		Raft:        n,
		ServerStats: ss,
		LeaderStats: stats.NewLeaderStats(strconv.Itoa(int(n.ID))),
		ErrorC:      make(chan error),
	}
	n.transport.Start()
	for i := range n.Peers {
		if i+1 != n.ID {
			n.transport.AddPeer(types.ID(i+1), []string{n.Peers[i]})
		}
	}

	go n.serverRaft()
	go n.serverChannels()
}

func (n *Node) serverChannels() {
	// heart beat tick
	ticker := time.NewTicker(heartBeatDuration)
	defer ticker.Stop()

	// send proposals over raft
	go func() {
		for n.proposeC != nil {
			select {
			case prop, ok := <-n.proposeC:
				if !ok {
					n.proposeC = nil
				} else {
					n.node.Propose(context.TODO(), []byte(prop))
				}
			}
		}
		close(n.stopC)
	}()

	for {
		select {
		case <-ticker.C:
			n.node.Tick()
		}
	}
}

func (n *Node) serverRaft() {

}

// PublishEntries writes commited log entries to commit channel
// and returns whether entries could be published.
func (n *Node) PublishEntries(ents []raftpb.Entry) bool {
	for i := range ents {
		switch ents[i].Type {
		case raftpb.EntryNormal:
			if len(ents[i].Data) == 0 {
				break
			}
			s := string(ents[i].Data)
			select {
			case n.commitC <- &s:
			case <-n.stopC:
				return false
			}
		// TODO:
		case raftpb.EntryConfChange:
		}
		// Alter commit, update appliedIndex.
		n.appliedIndex = ents[i].Index

		// Special nil commit to signal replay has finished.
		if ents[i].Index == n.lastIndex {
			select {
			case n.commitC <- nil:
			case <-n.stopC:
				return false
			}
		}
	}

	return true
}

// Process :
func (n *Node) Process(ctx context.Context, m raftpb.Message) error {
	return n.node.Step(ctx, m)
}

// IsIDRemoved :
func (n *Node) IsIDRemoved(id uint64) bool { return false }

// ReportUnreachable :
func (n *Node) ReportUnreachable(id uint64) {}

// ReportSnapshot :
func (n *Node) ReportSnapshot(id uint64, status raft.SnapshotStatus) {}

// New return a Raft Node
func New(conf *Conf) *Node {
	proposeC := make(chan string)
	commitC := make(chan *string)
	errorC := make(chan error)
	stopC := make(chan struct{})
	backend := backend.NewInfluxBackend()

	return &Node{
		ID:       conf.ID,
		Peers:    conf.Peers,
		Join:     conf.Join,
		proposeC: proposeC,
		commitC:  commitC,
		errorC:   errorC,
		stopC:    stopC,
		Backend:  backend,
		storage:  raft.NewMemoryStorage(),
	}
}
