package main

import (
	"flag"
	"strings"

	"io/ioutil"

	"github.com/lwhile/influxcap/node"
	"github.com/lwhile/influxcap/service"
	"github.com/lwhile/log"
	"gopkg.in/yaml.v2"
)

var (
	idFlag      = flag.Int("id", 0, "node id")
	confFlag    = flag.String("conf", "conf.yml", "config file path")
	joinFlag    = flag.Bool("join", false, "join a influxcap cluster")
	portFlag    = flag.String("port", "4928", "http listen port")
	clusterFlag = flag.String("cluster", "", "cluster address to join")
)

type config struct {
	Addr        string `yaml:"addr"`
	Name        string `yaml:"name"`
	InfluxdbURL string `yaml:"influxdb_url"`
}

func parseConfig(path string) (*config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf := config{}
	err = yaml.Unmarshal(b, &conf)
	return &conf, err
}

func main() {
	flag.Parse()

	if *confFlag == "" {
		log.Fatal("must specific a path of config file")
	}

	// parse the running config
	conf, err := parseConfig(*confFlag)
	if err != nil {
		log.Fatalf("influxcap: Failed to parse a config file (%v)", err)
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

	serverConf := service.ServerConf{
		Port: conf.Addr,
	}
	server := service.NewServer(&serverConf)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
