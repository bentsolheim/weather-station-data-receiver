package app

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
)

func CreateGinEngine(dataLogController DataLogController) *gin.Engine {
	r := gin.Default()
	r.Use(RequestLoggerMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/logger/:id/readings", dataLogController.GetReadings)
		v1.POST("/logger/:id/readings", dataLogController.PostReadings)
		v1.GET("/logger/:id/debug", dataLogController.GetDebug)
		v1.POST("/logger/:id/debug", dataLogController.PostDebug)
	}

	return r
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ := ioutil.ReadAll(tee)
			c.Request.Body = ioutil.NopCloser(&buf)
			log.Print(string(body))
		}
		c.Next()
	}
}
