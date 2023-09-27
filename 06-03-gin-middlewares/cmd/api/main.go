package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/06-03-gin-middlewares/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
