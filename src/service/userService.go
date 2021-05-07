package service

import (
	"chargeMsGo/src/database"
	"chargeMsGo/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var result models.Result

func UserLogin(ctx *gin.Context) models.Result {
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
		user := models.User{}
		ctx.ShouldBind(&user)
		tx := database.Db.Raw("SELECT id,name,phoneNumber,carid,cellid,avatar FROM user WHERE username = ? and password = ?", user.Username,user.Password).Scan(&user)
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

func UserRegister(ctx *gin.Context) models.Result {
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
	user := models.User{}
	ctx.ShouldBind(&user)
	if len(user.Avatar) == 0 {
		user.Avatar = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
	}
	tx := database.Db.Exec("INSERT INTO user(name,username,password,phoneNumber,avatar) VALUES (?,?,?,?,?);",user.Name,user.Username,user.Password,user.PhoneNumber,user.Avatar)
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
