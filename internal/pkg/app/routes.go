package app

import (
	"github.com/gin-gonic/gin"
)

func CreateGinEngine(dataLogController DataLogController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/logger/:id/readings", dataLogController.GetReadings)
		v1.POST("/logger/:id/readings", dataLogController.PostReadings)
		v1.POST("/logger/:id/debug", dataLogController.PostDebug)
	}

	return r
}
