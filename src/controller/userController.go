package controller

import (
	. "chargeMsGo/src/models"
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var result Result

func UserLogin(ctx *gin.Context)  {
	result = service.UserLogin(ctx)
	ctx.JSON(http.StatusOK, result)
}

func UserRegister(ctx *gin.Context)  {
	result = service.UserRegister(ctx)
	ctx.JSON(http.StatusOK, result)
}
