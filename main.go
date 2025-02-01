package main

import (
	"flag"
	"log/slog"
	"os"

	m "github.com/majermarci/go-mock/mock"
)

var configPath = flag.String("c", "config.yaml", "Path to the config file (default: config.yaml)")

func main() {
	flag.Parse()

	conf, err := m.LoadConfig(*configPath)
	if err != nil {
		slog.Error("Cannot load config", "details", err.Error())
		os.Exit(1)
	}

	m.Serve(conf)
}
