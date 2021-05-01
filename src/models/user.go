package models

type User struct {
	Id       	int    `form:"id"      		json:"id"       	gorm:"column:id;primary_key"`
	Name    	string `form:"name" 		json:"name"  	 	gorm:"column:name"`
	Username 	string `form:"username" 	json:"username"  	gorm:"column:username"`
	Password 	string `form:"password" 	json:"password"  	gorm:"column:passowrd"`
	PhoneNumber string `form:"phoneNumber"  json:"phoneNumber"  gorm:"column:phoneNumber"`
	Carid		int	   `form:"carid"  		json:"carid"  		gorm:"column:carid"`
	Cellid		int	   `form:"cellid"  		json:"cellid"  		gorm:"column:cellid"`
	Avatar   	string `form:"avatar"   	json:"avatar"   	gorm:"column:avatar"`
}

//type User struct {
//	Id       	int    `form:"id"      		json:"id"       	gorm:"column:id;primary_key"`
//	Name    	string `form:"name" 		json:"name"  	 	gorm:"column:name"`
//	Username 	string `form:"username" 	json:"username"  	gorm:"column:username"`
//	Password 	string `form:"password" 	json:"password"  	gorm:"column:passowrd"`
//	PhoneNumber string `form:"phoneNumber"  json:"phoneNumber"  gorm:"column:phoneNumber"`
//	CarId		int	   `form:"carId"  		json:"carId"  		gorm:"column:carId"`
//	CellId		int	   `form:"cellId"  		json:"cellId"  		gorm:"column:cellId"`
//	Avatar   	string `form:"avatar"   	json:"avatar"   	gorm:"column:avatar"`
//}
