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
		tx := database.Db.Raw("SELECT id,name,phoneNumber,carId,cellId FROM user WHERE username = ? and password = ?", user.Username,user.Password).Scan(&user)
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