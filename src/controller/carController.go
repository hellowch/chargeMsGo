package controller

import (
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserCar(ctx *gin.Context)  {
	result = service.GetUserCar(ctx)
	ctx.JSON(http.StatusOK, result)
}

func SetUserCar(ctx *gin.Context)  {
	result = service.SetUserCar(ctx)
	ctx.JSON(http.StatusOK, result)
}
