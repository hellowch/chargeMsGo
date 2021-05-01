package controller

import (
	. "chargeMsGo/src/models"
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var result Result

//func SelectUserById(ctx *gin.Context) {
//	result = service.SelectUserById(ctx)
//	ctx.JSON(http.StatusOK, result)
//}
//
//func GormInsertData(ctx *gin.Context) {
//	result = service.GormInsertData(ctx)
//	ctx.JSON(http.StatusOK, result)
//}

func UserLogin(ctx *gin.Context)  {
	result = service.UserLogin(ctx)
	ctx.JSON(http.StatusOK, result)
}
