package app

import (
	"github.com/bentsolheim/go-app-utils/db"
	"github.com/bentsolheim/go-app-utils/utils"
)

type AppConfig struct {
	Db         db.DbConfig
	ServerPort string
}

func ReadAppConfig() AppConfig {
	dbc := db.ReadDbConfig(db.DbConfig{
		User:     "nta_consultants",
		Password: "nta_consultants",
		Host:     "localhost",
		Port:     "3306",
		Name:     "nta_consultants",
	})
	return AppConfig{
		dbc,
		utils.GetEnvOrDefault("SERVER_PORT", "8080"),
	}
}
