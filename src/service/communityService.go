package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"fmt"
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
	resCommunity := make([]models.ResCommunity,0)
	userid := ctx.DefaultQuery("userid", "nil")
	if userid == "nil" {
		tx := database.Db.Debug().Raw("SELECT community.id,community.userid,community.details,community.detailsid,user.name,user.avatar FROM user,community WHERE user.id = community.userid ORDER BY community.id DESC").Find(&resCommunity)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}else if userid != "nil" {
		tx := database.Db.Debug().Raw("SELECT id,userid,details,detailsid FROM community WHERE userid = ?", userid).Find(&resCommunity)
		if tx.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = tx.Error
			return result
		}
	}

	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = resCommunity
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

func GetCarCommunity(ctx *gin.Context) models.Result {
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
	resCommunity := make([]models.ResCommunity,0)
	userCar := ctx.DefaultQuery("userCar", "nil")
	car := make([]models.Car,0)
	tx := database.Db.Debug().Where(fmt.Sprintf("carbrand LIKE %q",(userCar+"%"))).Find(&car)
	userMap := make(map[int]int)
	for i:=0;i<len(car);i++ {
		userMap[car[i].Userid] = 1
	}
	useridSlice := make([]int,0)
	for k,_ := range userMap{
		useridSlice = append(useridSlice, k)
	}
	tx = database.Db.Debug().Raw("SELECT community.id,community.userid,community.details,community.detailsid,user.name,user.avatar FROM community,user where user.id = community.userid and userid IN (?) ORDER BY id DESC", useridSlice).Scan(&resCommunity)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}

	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = resCommunity
	return result
}
