package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCommunity(ctx *gin.Context) models.Result {
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
	community := make([]models.Community,0)
	userid := ctx.DefaultQuery("userid", "nil")
	if userid == "nil" {
		tx := database.Db.Debug().Raw("SELECT id,userid,details,detailsid FROM community").Find(&community)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}else if userid != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,userid,details,detailsid FROM community WHERE userid = ?", userid).Find(&community)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}

	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = community
	return result
}

func SetCommunity(ctx *gin.Context) models.Result  {
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
	community := models.Community{}
	ctx.ShouldBind(&community)
	tx := database.Db.Exec("INSERT INTO community(userid,details,detailsid) VALUES (?,?,?);",community.Userid,community.Details,community.Detailsid)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "插入成功"
	result.Data = tx.RowsAffected
	return result
}
