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
	user.POST("/register",controller.UserRegister)

	car := Engine.Group("/carApi")
	car.GET("/getUserCar",controller.GetUserCar)
	car.POST("/setUserCar", controller.SetUserCar)

	charger := Engine.Group("/charger")
	charger.GET("/getCharger",controller.GetCharger)
	charger.GET("/getChargerDetails",controller.GetChargerDetails)
	charger.GET("/getChargerOrder",controller.GetChargerOrder)
	charger.GET("/getSumOrder",controller.GetSumOrder)

	Engine.Run(":3000")

}
