package main

import (
	"log"
)

func main() {
	conf, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	serve(conf)
}
