package service

import "net/http"

const (
	// post
	joinAPI = "/join"
)

// JoinRequest :
type JoinRequest struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func serveJoin(w http.ResponseWriter, r *http.Request) {

}
