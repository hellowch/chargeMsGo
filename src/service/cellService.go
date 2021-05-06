package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCell(ctx *gin.Context) models.Result {
	//捕获异常
	defer func() {
		err := recover()
		if err != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误1"
			result.Data = err
			return
		}
	}()
	cell := models.Cell{}
	cellId := ctx.DefaultQuery("cellId", "nil")
	tx := database.Db.Debug().Raw("SELECT id,health,exengine,exair,exchip,exnav,exlamp,advice,electricity FROM cell WHERE id = ?", cellId).Scan(&cell)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = cell
	return result
}
