package main

import (
	"fmt"
	"github.com/diegodhdev/hexagonal-go-api/final/cmd/api/bootstrap"
	"log"
)

func main() {
	fmt.Println("main.go")
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
