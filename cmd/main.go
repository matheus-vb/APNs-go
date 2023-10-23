package main

import (
	"github.com/joho/godotenv"
	"go-server/api"
	"go-server/internal"
	"log"
	"os"
)

func main() {
	log.Println("Setting up server...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := api.App{
		Provider: internal.Provider{
			KeyId:   os.Getenv("KEY_ID"),
			TeamId:  os.Getenv("TEAM_ID"),
			KeyFile: "./auth-key.p8",
			Topic:   os.Getenv("BUNDLE_ID"),
			Client:  nil,
		},
	}

	app.SetupServer()
}
