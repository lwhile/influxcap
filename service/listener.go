package service

import (
	"errors"
	"net"
	"time"
)

// Listener :
type Listener struct {
	*net.TCPListener
	stopC <-chan struct{}
}

// NewListener :
func NewListener(addr string, stopC <-chan struct{}) (*Listener, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Listener{
		ln.(*net.TCPListener),
		stopC,
	}, nil
}

func (ln *Listener) Accept() (c net.Conn, err error) {
	connc := make(chan *net.TCPConn, 1)
	errC := make(chan error, 1)
	go func() {
		tc, err := ln.AcceptTCP()
		if err != nil {
			errC <- err
			return
		}
		connc <- tc
	}()
	select {
	case <-ln.stopC:
		return nil, errors.New("server stopped")
	case err := <-errC:
		return nil, err
	case tc := <-connc:
		tc.SetKeepAlive(true)
		tc.SetKeepAlivePeriod(3 * time.Minute)
		return tc, nil
	}
}
