package raft

import "github.com/lwhile/influxcap/backend"

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

// Node is a influxcap instance,
// which be implemented as a raft node(Raft interface)
type Node struct {
	ID     string   `json:"id"`
	Status Status   `json:"status"`
	Peers  []string `json:"peers"`

	proposeC <-chan string
	commitC  chan<- string
	errorC   chan<- error

	stopC   chan struct{}
	Backend *backend.Influxdb
}

// NodeConf is config of a Node
type NodeConf struct {
}

// NewNode return a Raft Node
func NewNode(ID string) *Node {
	proposeC := make(chan string)
	commitC := make(chan string)
	errorC := make(chan error)
	stopC := make(chan struct{})
	backend := backend.NewInfluxBackend()

	return &Node{
		ID:       ID,
		proposeC: proposeC,
		commitC:  commitC,
		errorC:   errorC,
		stopC:    stopC,
		Backend:  backend,
	}
}
