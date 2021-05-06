package controller

import (
	"chargeMsGo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCell(ctx *gin.Context)  {
	result = service.GetCell(ctx)
	ctx.JSON(http.StatusOK, result)
}
