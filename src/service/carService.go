package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GatUserCar(ctx *gin.Context) models.Result {
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

func SetUserCar(ctx *gin.Context) models.Result {
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
	ctx.ShouldBind(&car)
	fmt.Println(car)
	tx := database.Db.Exec("update car set carbrand=?,carmodel=?,buytime=? WHERE id=?",car.Carbrand,car.Carmodel,car.Buytime.Format("2006-01-02"),car.Id)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "修改成功"
	result.Data = tx.RowsAffected
	return result
}
