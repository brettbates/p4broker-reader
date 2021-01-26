package main

import (
	"log"
	"os"

	reader "github.com/brettbates/p4broker-reader/reader"
)

func main() {
	res, err := reader.Read(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", res)
}
