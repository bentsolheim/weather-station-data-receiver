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

	dataLogController := app.DataLogController{DataLogService: app.NewDataLogService()}

	router := app.CreateGinEngine(dataLogController)
	return router.Run(fmt.Sprintf(":%s", c.ServerPort))
}
