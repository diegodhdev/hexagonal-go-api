package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/01-03-architectured-healthcheck/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
