package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/08-02-build-and-deploy/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
