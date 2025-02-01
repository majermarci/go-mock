package main

import (
	"flag"
	"log/slog"
	"os"
)

var configPath = flag.String("c", "config.yaml", "Path to the config file (default: config.yaml)")

func main() {
	flag.Parse()

	conf, err := loadConfig(*configPath)
	if err != nil {
		slog.Error("Cannot load config", "details", err.Error())
		os.Exit(1)
	}

	serve(conf)
}
