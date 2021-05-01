package routers

import (
	"chargeMsGo/src/controller"
	"chargeMsGo/src/middleware"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func InitRouter() {
	Engine := gin.Default()
	Engine.Use(middleware.RequestInfos(),middleware.Cors())

	//user 路由组
	user := Engine.Group("/api")
	user.POST("/login",controller.UserLogin)

	Engine.Run(":3000")

}
