package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/05-02-timeouts/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
