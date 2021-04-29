package controller

import (
	. "chargeMsGo/src/database"
	. "chargeMsGo/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var result Result

func SelectUserById(context *gin.Context) {
	//捕获异常
	defer func() {
		err := recover()
		if err != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误1"
			result.Data = err
			context.JSON(http.StatusBadRequest, result)
		}
	}()
	//获取问号后面请求的数据,获取不到就返回hello默认值
	id := context.DefaultQuery("id", "null")
	fmt.Println(id)
	user := User{}
	tx := Db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "错误"
		result.Data = tx.Error
		context.JSON(http.StatusOK, result)
		return
	}
	result.Code = http.StatusOK
	result.Message = "查询成功"
	result.Data = user
	context.JSON(http.StatusOK, result)
}

func GormInsertData(c *gin.Context) {
	//=============捕获异常============
	defer func() {
		err := recover()
		if err != nil {
			result.Code = http.StatusBadRequest
			result.Message = "错误"
			result.Data = err
			c.JSON(http.StatusBadRequest, result)
		}
	}()
	//============
	var user User
	err := c.Bind(&user)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "参数错误"
		result.Data = err
		c.JSON(http.StatusOK, result)
		return
	}
	fmt.Println(user)
	tx := Db.Create(&user)
	if tx.RowsAffected > 0 {
		result.Code = http.StatusOK
		result.Message = "写入成功"
		result.Data = "OK"
		c.JSON(http.StatusOK, result)
		return
	}
	fmt.Printf("insert failed, err:%v\n", err)
	result.Code = http.StatusBadRequest
	result.Message = "写入失败"
	result.Data = tx
	c.JSON(http.StatusOK, result)
	fmt.Println(tx) //打印结果
}
