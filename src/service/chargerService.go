package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCharger(ctx *gin.Context) models.Result {
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
	charger := models.Charger{}
	id := ctx.DefaultQuery("id", "nil")
	tx := database.Db.Debug().Raw("SELECT id,address,distance,usetype,details FROM charger WHERE id = ?", id).Scan(&charger)
	fmt.Println(charger)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = charger
	return result
}

func GetChargerDetails(ctx *gin.Context) models.Result {
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
	chargerDetails := models.ChargerDetails{}
	chargerId := ctx.DefaultQuery("chargerId", "nil")
	tx := database.Db.Debug().Raw("SELECT id,chargerid,details,time FROM chargerDetails WHERE chargerid = ?", chargerId).Scan(&chargerDetails)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = chargerDetails
	return result
}

func GetChargerOrder(ctx *gin.Context) models.Result {
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
	chargerOrder := models.ChargerOrder{}
	userId := ctx.DefaultQuery("userId", "nil")
	chargerId := ctx.DefaultQuery("chargerId", "nil")

	if userId != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,chargerid,userid,amount,time,length FROM chargerOrder WHERE userid = ?", userId).Scan(&chargerOrder)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}else if chargerId != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,chargerid,userid,amount,time,length FROM chargerOrder WHERE chargerid = ?", chargerId).Scan(&chargerOrder)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = chargerOrder
	return result
}
