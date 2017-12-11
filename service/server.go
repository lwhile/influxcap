package service

import "net/http"

// ServerConf contain config of a http server
type ServerConf struct {
	Addr string
}

// Server http request
type Server struct {
	conf      *ServerConf
	server    http.Server
	httpstopC chan struct{}
}

// NewServer return a Server instance
func NewServer(conf *ServerConf) *Server {
	return &Server{
		conf: conf,
	}
}

// Start a http server
// this method will be blocked
func (s *Server) Start() error {
	s.server.Addr = s.conf.Addr
	return s.server.ListenAndServe()
}

// Stop the http server
// TODO: implement
func (s *Server) Stop() error {
	return nil
}
