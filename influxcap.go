package main

import (
	"flag"

	"github.com/lwhile/influxcap/node"
	"github.com/lwhile/influxcap/service"
	"github.com/lwhile/log"
)

var (
	idFlag   = flag.Int("id", 0, "node id")
	confFlag = flag.String("conf", "conf.yml", "config file path")
	joinFlag = flag.Bool("join", false, "join a influxcap cluster")
	portFlag = flag.String("port", "4928", "http listen port")
)

func main() {
	flag.Parse()
	if *idFlag == 0 {
		log.Fatal("Node ID must no be zero")
	}
	nodeConf := node.Conf{
		ID:   *idFlag,
		Join: *joinFlag,
	}
	node := node.New(&nodeConf)
	node.Start()

	serverConf := service.ServerConf{
		Port: *portFlag,
	}
	server := service.NewServer(&serverConf)

	server.Start()
}
