package app

import (
	"github.com/bentsolheim/go-app-utils/utils"
)

type AppConfig struct {
	ServerPort string
}

func ReadAppConfig() AppConfig {
	return AppConfig{
		utils.GetEnvOrDefault("SERVER_PORT", "8080"),
	}
}
