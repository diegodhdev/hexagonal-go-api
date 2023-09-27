package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/03-02-repository-test/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
