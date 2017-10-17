package base

// DBNode container actions about each node
type DBNode interface {
	Start() error
	Stop() error
	Status() *NodeStatus
}

// DBCluster container actions about a cluster
type DBCluster interface {
	Join() error
	Quit() error
	Status() *ClusterStatus
}

// ClusterStatus maintain a cluster status
type ClusterStatus struct {
}

// InfluxCluster provides infrastructure a cluster instance
// it implements DBCluster interface
type InfluxCluster struct {
}

// InfluxNode provides infrastructure of a influxdb instance
// it implements DBNode interface
type InfluxNode struct {
	ID   uint64
	Name string
}

// NodeStatus maintain a node status
type NodeStatus struct {
}
