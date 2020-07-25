package app

import (
	"github.com/bentsolheim/go-app-utils/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Readings struct {
	Readings []SensorReading
}

type DataLogController struct {
	DataLogService *DataLogService
}

func (c *DataLogController) GetReadings(ctx *gin.Context) {

	loggerId := ctx.Param("id")
	readings, err := c.DataLogService.FindLatestReadings(loggerId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rest.WrapResponse(nil, err))
		return
	}
	ctx.JSON(http.StatusOK, rest.WrapResponse(readings, err))
}

func (c *DataLogController) PostReadings(ctx *gin.Context) {

	var data Readings
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, rest.WrapResponse(nil, err))
		return
	}
	loggerId := ctx.Param("id")
	for _, reading := range data.Readings {
		if err := c.DataLogService.RegisterData(loggerId, reading); err != nil {
			ctx.JSON(http.StatusInternalServerError, rest.WrapResponse(nil, err))
			return
		}
	}
	ctx.Status(http.StatusNoContent)
}

func (c *DataLogController) PostDebug(ctx *gin.Context) {

	var debug Debug
	if err := ctx.ShouldBindJSON(&debug); err != nil {
		ctx.JSON(http.StatusBadRequest, rest.WrapResponse(nil, err))
		return
	}
	loggerId := ctx.Param("id")
	if err := c.DataLogService.RegisterDebug(loggerId, debug); err != nil {
		ctx.JSON(http.StatusInternalServerError, rest.WrapResponse(nil, err))
		return
	}
	ctx.Status(http.StatusNoContent)
}
