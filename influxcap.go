package main

import (
	"flag"

	_ "github.com/coreos/etcd/raft"
)

var (
	confFlag = flag.String("conf", "conf.yml", "config file path")
	joinFlag = flag.Bool("join", false, "join a influxcap cluster")
	portFlag = flag.String("port", "4928", "http listen port")
)

func main() {
	flag.Parse()

}
