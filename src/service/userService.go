package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var result models.Result

func SelectUserById(ctx *gin.Context) models.Result {
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
	//获取问号后面请求的数据,获取不到就返回hello默认值
	id := ctx.DefaultQuery("id", "null")
	fmt.Println(id)
	user := models.User{}
	tx := database.Db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		return result
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = user
	return result
}

func GormInsertData(ctx *gin.Context) models.Result {
	//=============捕获异常============
	defer func() {
		err := recover()
		if err != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = err
			return
		}
	}()
	//============
	var user models.User
	err := ctx.Bind(&user)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "参数错误"
		result.Data = err
		return result
	}
	fmt.Println(user)
	tx := database.Db.Create(&user)
	if tx.RowsAffected > 0 {
		result.Code = http.StatusOK
		result.Message = "写入成功"
		result.Data = "OK"
		return result
	}
	fmt.Printf("insert failed, err:%v\n", err)
	result.Code = http.StatusBadRequest
	result.Message = "写入失败"
	result.Data = tx
	return result
}
