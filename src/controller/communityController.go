package controller

import (
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCommunity(ctx *gin.Context)  {
	result = service.GetCommunity(ctx)
	ctx.JSON(http.StatusOK, result)
}

func SetCommunity(ctx *gin.Context)  {
	result = service.SetCommunity(ctx)
	ctx.JSON(http.StatusOK, result)
}

func GetCarCommunity(ctx *gin.Context)  {
	result = service.GetCarCommunity(ctx)
	ctx.JSON(http.StatusOK, result)
}
