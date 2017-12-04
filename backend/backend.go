package backend

const (
	// CONNECT status
	CONNECT Status = "connecting"
	// UNCONNECT status
	UNCONNECT Status = "unconnecting"
)

// Status type
type Status string

// Backend interface
type Backend interface {
}

// Influxdb as a influxdb server
// it implement Backend interface
type Influxdb struct {
	URL    string `json:"url"`
	Status Status `json:"status"`
}

// NewInfluxBackend will return a influxdb db backend
func NewInfluxBackend() *Influxdb {
	return nil
}
