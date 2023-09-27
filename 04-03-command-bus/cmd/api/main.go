package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/04-03-command-bus/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
