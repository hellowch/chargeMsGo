package routers

import (
	"chargeMsGo/src/controller"
	"chargeMsGo/src/middleware"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func InitRouter() {
	Engine := gin.Default()
	Engine.Use(middleware.RequestInfos())

	//user 路由组
	user := Engine.Group("/user")
	user.GET("/selectUserById", controller.SelectUserById)
	user.POST("/gormInsertData", controller.GormInsertData)

	Engine.Run(":3000")

}
