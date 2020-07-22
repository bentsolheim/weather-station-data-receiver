package app

import (
	"github.com/gin-gonic/gin"
)

func CreateGinEngine(dataLogController DataLogController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/data-log-request", dataLogController.LogData)
	}

	return r
}
