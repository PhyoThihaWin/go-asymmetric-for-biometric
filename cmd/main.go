package main

import (
	"log"
	"os"

	"pthw.com/asymmetric-for-biometric/config"
	"pthw.com/asymmetric-for-biometric/server"
)

func main() {
	if err := config.InitGoDotEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := server.NewApp()
	if err := app.Run(os.Getenv("WEB_HOST")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
