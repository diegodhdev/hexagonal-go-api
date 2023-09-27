package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/05-01-graceful-shutdown/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
