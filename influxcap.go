package main

import (
	"flag"
	"strings"

	"github.com/lwhile/influxcap/node"
	"github.com/lwhile/log"
)

var (
	idFlag      = flag.Int("id", 0, "node id")
	confFlag    = flag.String("conf", "conf.yml", "config file path")
	joinFlag    = flag.Bool("join", false, "join a influxcap cluster")
	portFlag    = flag.String("port", "4928", "http listen port")
	clusterFlag = flag.String("cluster", "", "cluster address to join")
)

func main() {
	flag.Parse()

	if *idFlag == 0 {
		log.Fatal("node ID must no be zero")
	}

	if *joinFlag && *clusterFlag == "" {
		log.Fatal("must specific a cluster address to join")
	}

	nodeConf := node.Conf{
		ID:    *idFlag,
		Join:  *joinFlag,
		Peers: strings.Split(*clusterFlag, ","),
	}
	node := node.New(&nodeConf)
	if *joinFlag {
		node.Peers = []string{*clusterFlag}
	}
	node.Start()

	// serverConf := service.ServerConf{
	// 	Port: *portFlag,
	// }
	// server := service.NewServer(&serverConf)

	// if err := server.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	select {}
}
