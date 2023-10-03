package main

import (
	"github.com/diegodhdev/hexagonal-go-api/final/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
