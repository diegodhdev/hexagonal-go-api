package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/08-01-reading-env-variables/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
