package app

import (
	"github.com/gin-gonic/gin"
)

func CreateGinEngine(consultantsController DataPayloadController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	reports := v1.Group("/reports")
	{
		reports.GET("/most-recent-temp", consultantsController.MostRecentTemp)
	}

	return r
}
