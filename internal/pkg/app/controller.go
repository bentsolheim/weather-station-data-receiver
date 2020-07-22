package app

import (
	"github.com/bentsolheim/go-app-utils/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	LoggerId   string
	SensorName string
	Value      float32
}

type DataLogController struct {
	DataLogService *DataLogService
}

func (c *DataLogController) LogData(ctx *gin.Context) {
	var data Data
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, rest.WrapResponse(nil, err))
		return
	}
	if err := c.DataLogService.RegisterData(data.LoggerId, data.SensorName, data.Value); err != nil {
		ctx.JSON(http.StatusInternalServerError, rest.WrapResponse(nil, err))
		return
	}
	ctx.Status(http.StatusNoContent)
}
