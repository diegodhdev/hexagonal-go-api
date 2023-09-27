package main

import (
	"log"

	"github.com/diegodhdev/hexagonal-go-api/02-01-post-course-endpoint/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
