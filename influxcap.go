package main

import (
	"flag"

	_ "github.com/coreos/etcd/raft"
	"github.com/lwhile/influxcap/service"
)

var (
	confFlag = flag.String("conf", "conf.yml", "config file path")
	joinFlag = flag.Bool("join", false, "join a influxcap cluster")
	portFlag = flag.String("port", "4928", "http listen port")
)

func main() {
	flag.Parse()

	serverConf := service.ServerConf{
		Port: *portFlag,
	}
	server := service.NewServer(&serverConf)

	server.Start()
}
