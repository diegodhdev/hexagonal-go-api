package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/07-01-publishing-domain-events/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
