package main

import (
	"flag"
	"log"
)

var configPath = flag.String("c", "config.yaml", "Path to the config file (default: config.yaml)")

func main() {
	flag.Parse()

	conf, err := loadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	serve(conf)
}
