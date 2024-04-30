package main

import (
	"log"

	"github.com/robowealth-mutual-fund/blueprint-service/internals/container"
)

func main() {
	server, err := container.New()
	if err != nil {
		log.Panic(err)
	}

	if err = server.Start(); err != nil {
		log.Panic(err)
	}

}
