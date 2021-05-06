package controller

import (
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCharger(ctx *gin.Context)  {
	result = service.GetCharger(ctx)
	ctx.JSON(http.StatusOK, result)
}

func GetChargerDetails(ctx *gin.Context)  {
	result = service.GetChargerDetails(ctx)
	ctx.JSON(http.StatusOK, result)
}

func GetChargerOrder(ctx *gin.Context)  {
	result = service.GetChargerOrder(ctx)
	ctx.JSON(http.StatusOK, result)
}

func GetSumOrder(ctx *gin.Context)  {
	result = service.GetSumOrder(ctx)
	ctx.JSON(http.StatusOK, result)
}
