package main

import (
	"fmt"
	"github.com/bentsolheim/weather-station-data-receiver/internal/pkg/app"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	c := app.ReadAppConfig()

	consultantsController := app.DataPayloadController{DataPayloadService: &app.DataPayloadService{}}

	router := app.CreateGinEngine(consultantsController)
	return router.Run(fmt.Sprintf(":%s", c.ServerPort))
}
