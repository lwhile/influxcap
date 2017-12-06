package raft

import "github.com/lwhile/influxcap/backend"
import "github.com/coreos/etcd/raft"

const (
	// RUNNING status
	RUNNING Status = "running"

	// DOWNING status
	DOWNING Status = "downing"
)

// Raft interface
type Raft interface {
}

// Status type
type Status string

// NodeConf is config of a Node
type NodeConf struct {
	ID    string
	Peers []string
	Join  bool
}

// Node is a influxcap instance,
// which be implemented as a raft node(Raft interface)
type Node struct {
	ID     string   `json:"id"`
	Status Status   `json:"status"`
	Join   bool     `json:"join"`
	Peers  []string `json:"peers"`

	proposeC <-chan string
	commitC  chan<- string
	errorC   chan<- error

	stopC   chan struct{}
	Backend *backend.Influxdb

	// Field below is relating to raft library

	// etcd raft node interface for the commit/error channel
	node raft.Node

	storage *raft.MemoryStorage
}

// Start a raft node
func (n *Node) Start() error {
	return nil
}

// NewNode return a Raft Node
func NewNode(conf *NodeConf) *Node {
	proposeC := make(chan string)
	commitC := make(chan string)
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
	}
}
