package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/06-02-time-parse-in-go/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
