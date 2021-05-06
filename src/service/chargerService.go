package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	charger := make([]models.Charger,0)
	id := ctx.DefaultQuery("id", "nil")
	tx := database.Db.Debug().Raw("SELECT id,address,distance,usetype,details FROM charger WHERE id = ?", id).Find(&charger)
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
	chargerDetails := make([]models.ChargerDetails,0)
	chargerId := ctx.DefaultQuery("chargerId", "nil")
	tx := database.Db.Debug().Raw("SELECT id,chargerid,details,time FROM charger_details WHERE chargerid = ?", chargerId).Find(&chargerDetails)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	resChargerDetails := make([]models.ResChargerDetails,len(chargerDetails))
	for i:=0;i<len(chargerDetails);i++ {
		resChargerDetails[i].Chargerid = chargerDetails[i].Chargerid
		resChargerDetails[i].Id = chargerDetails[i].Id
		resChargerDetails[i].Details = chargerDetails[i].Details
		resChargerDetails[i].Time = chargerDetails[i].Time.Format("2006-01-02")
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
	chargerOrder := make([]models.ChargerOrder,0)
	userId := ctx.DefaultQuery("userId", "nil")
	chargerId := ctx.DefaultQuery("chargerId", "nil")

	if userId != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,chargerid,userid,amount,time,length FROM charger_order WHERE userid = ? order by id desc", userId).Find(&chargerOrder)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}else if chargerId != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,chargerid,userid,amount,time,length FROM charger_order WHERE chargerid = ? order by id desc", chargerId).Find(&chargerOrder)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}
	resChargerOrder := make([]models.ResChargerOrder,len(chargerOrder))
	for i:=0;i<len(chargerOrder);i++ {
		resChargerOrder[i].Id = chargerOrder[i].Id
		resChargerOrder[i].Chargerid = chargerOrder[i].Chargerid
		resChargerOrder[i].Length = strconv.Itoa(chargerOrder[i].Length) + "分钟"
		resChargerOrder[i].Userid = chargerOrder[i].Userid
		resChargerOrder[i].Amount = chargerOrder[i].Amount + "元"
		resChargerOrder[i].Time = chargerOrder[i].Time.Format("2006-01-02")
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = resChargerOrder
	return result
}

func GetSumOrder(ctx *gin.Context) models.Result {
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
	chargerOrder := make([]models.ChargerOrder,0)
	userId := ctx.DefaultQuery("userId", "nil")
	chargerId := ctx.DefaultQuery("chargerId", "nil")
	useTime := ctx.DefaultQuery("useTime", "nil")

	if userId != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,chargerid,userid,amount,time,length FROM charger_order WHERE DATE_FORMAT(time,'%Y-%m') = ? and userid = ? order by id desc", useTime,userId).Find(&chargerOrder)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}else if chargerId != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,chargerid,userid,amount,time,length FROM charger_order WHERE DATE_FORMAT(time,'%Y-%m') = ? and chargerid = ? order by id desc", chargerId).Find(&chargerOrder)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}
	resChargerOrder := make([]models.ResChargerOrder,len(chargerOrder))
	sumAmount := 0
	for i:=0;i<len(chargerOrder);i++ {
		resChargerOrder[i].Id = chargerOrder[i].Id
		resChargerOrder[i].Chargerid = chargerOrder[i].Chargerid
		resChargerOrder[i].Length = strconv.Itoa(chargerOrder[i].Length) + "分钟"
		resChargerOrder[i].Userid = chargerOrder[i].Userid
		resChargerOrder[i].Amount = chargerOrder[i].Amount + "元"
		resChargerOrder[i].Time = chargerOrder[i].Time.Format("2006-01-02")
		intAmount,_ := strconv.Atoi(chargerOrder[i].Amount)
		sumAmount += intAmount
	}
	mapData := make(map[string]interface{})
	mapData["chargerOrder"] = resChargerOrder
	mapData["sumAmount"] = sumAmount
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = mapData
	return result
}
