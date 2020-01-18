package main

import (
	"github.com/backery/server"
	"log"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal("Server crashed!", err.Error())
	}
}