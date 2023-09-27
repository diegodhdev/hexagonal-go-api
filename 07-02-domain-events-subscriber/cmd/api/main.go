package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/07-02-domain-events-subscriber/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
