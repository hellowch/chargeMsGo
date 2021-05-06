package controller

import (
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GatUserCar(ctx *gin.Context)  {
	result = service.GatUserCar(ctx)
	ctx.JSON(http.StatusOK, result)
}

func SetUserCar(ctx *gin.Context)  {
	result = service.SetUserCar(ctx)
	ctx.JSON(http.StatusOK, result)
}
