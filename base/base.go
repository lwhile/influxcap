package base

// RaftNode container actions about each node
type RaftNode interface {
	Start() error
	Stop() error
	Status() *NodeStatus
}

// RaftCluster container actions about a cluster
type RaftCluster interface {
	Join() error
	Quit() error
	Status() *ClusterStatus
}

// InfluxNode provides infrastructure of a influxdb instance
// it implements RaftNode interface
type InfluxNode struct {
	ID      uint64
	Name    string
	Port    string
	Backend *InfluxBackend
	Status  *NodeStatus
}

// InfluxBackend means a backend instance of influxdb
type InfluxBackend struct {
	Host string
	Port string
}

// NewInfluxNode return a InfluxNode instance
func NewInfluxNode(c *Config) *InfluxNode {
	if c == nil {
		return nil
	}
	// TODO: generate a unique id to influxNode
	return &InfluxNode{
		Name: c.Name,
		Port: c.Port,
		Backend: &InfluxBackend{
			Host: c.InfluxHost,
			Port: c.InfluxPort,
		},
	}
}

// InfluxCluster provides infrastructure a cluster instance
// it implements RaftCluster interface
type InfluxCluster struct {
	// cluster name
	Name   string
	Status *ClusterStatus
	Nodes  []*InfluxNode
}

// ClusterStatus maintain a cluster status
type ClusterStatus struct {
	Status Status
}

// NodeStatus maintain a node status
type NodeStatus struct {
	Status Status
}

// Status type
type Status string

const (
	// RUNNING say the cluster or node is working now
	RUNNING Status = "running"

	// DOWNNING say  the cluster or one node was down
	DOWNNING Status = "downning"

	// INACTIVE saysthe cluster or one node was inactive
	INACTIVE Status = "inactive"
)
