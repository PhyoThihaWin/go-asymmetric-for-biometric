package main

import (
	"log"

	"github.com/spf13/viper"
	"pthw.com/asymmetric-for-biometric/config"
	"pthw.com/asymmetric-for-biometric/server"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
