package main

import (
	"log"

	"github.com/aethiopicuschan/chiilang/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
