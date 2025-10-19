package main

import (
	"flag"
	"log/slog"
	"os"
)

var configPath = flag.String("c", "config.yaml", "Path to the config file (default: config.yaml)")

func main() {
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	}))
	slog.SetDefault(logger)

	conf, err := loadConfig(*configPath)
	if err != nil {
		slog.Error("cannot load config", "details", err.Error())
		os.Exit(1)
	}

	serve(conf)
}
