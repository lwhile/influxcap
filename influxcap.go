package main

import (
	"flag"
	"os"

	"io/ioutil"

	"github.com/levigross/grequests"
	"github.com/lwhile/influxcap/node"
	"github.com/lwhile/influxcap/service"
	"github.com/lwhile/log"
	"gopkg.in/yaml.v2"
)

var (
	confFlag = flag.String("conf", "conf.yml", "config file path")
	joinFlag = flag.String("join", "", "specific a cluster node to join")
)

type config struct {
	ID          uint64 `yaml:"id"`
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

func joinCluster(url string, id uint64, addr string) error {
	joinData := service.JoinRequest{
		Addr: addr,
		ID:   id,
	}
	resp, err := grequests.Post(url, &grequests.RequestOptions{JSON: joinData})
	if err != nil {
		return err
	}
	log.Infof("influxcap: join cluster %s %s", url, resp.String())
	return nil
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
		ID: conf.ID,
	}

	// join a culster
	if *joinFlag != "" {
		err := joinCluster(*joinFlag, conf.ID, conf.Addr)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	node := node.New(&nodeConf)
	node.Start()

	serverConf := service.ServerConf{
	//Port: conf.Addr,
	}
	server := service.NewServer(&serverConf)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
