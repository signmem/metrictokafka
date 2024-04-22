package main

import (
	"flag"
	"fmt"
	"github.com/signmem/metrictokafka/http"
	"github.com/signmem/metrictokafka/tools"
	"github.com/signmem/metrictokafka/g"
	_ "net/http/pprof"
	"os"
	"sync"
)


func init() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		version := g.Version
		fmt.Printf("%s", version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.Logger = g.InitLog()

	if g.Config().Kafka.Enable {
		tools.KafkaServer = g.Config().Kafka.Servers
		tools.Topic = g.Config().Kafka.Topic
	}

}

func main() {



	ch := make(chan *tools.MItem, 20)
	MItemList := tools.MItemList{}

	go MItemList.Start(ch, 20)

	var wg sync.WaitGroup

	go tools.GenHealthMetric(ch, &wg, 1)

	go http.Start()

	select {}
}
