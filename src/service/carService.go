package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CarUserCar(ctx *gin.Context) models.Result {
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
	car := models.Car{}
	carId := ctx.DefaultQuery("carId", "nil")
	tx := database.Db.Debug().Raw("SELECT id,carbrand,carmodel,buytime FROM car WHERE id = ?", carId).Scan(&car)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = car
	return result
}
