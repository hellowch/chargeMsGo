package controller

import (
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CarUserCar(ctx *gin.Context)  {
	result = service.CarUserCar(ctx)
	ctx.JSON(http.StatusOK, result)
}
