package main

import (
	"go-server/api"
	"log"
)

func main() {
	log.Println("Setting up server...")

	app := api.App{}

	app.SetupServer()
}
