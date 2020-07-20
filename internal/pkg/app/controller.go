package app

import (
	"github.com/bentsolheim/go-app-utils/rest"
	"github.com/gin-gonic/gin"
)

type DataPayloadController struct {
	DataPayloadService *DataPayloadService
}

func (c *DataPayloadController) MostRecentTemp(ctx *gin.Context) {
	temp, err := c.DataPayloadService.GetMostRecentTemp()
	ctx.JSON(200, rest.WrapResponse(temp, err))
}
