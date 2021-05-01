package models

import "time"

type Car struct {
	Id       	int        `form:"id"      		json:"id"       	gorm:"column:id;primary_key"`
	Carbrand    string     `form:"carbrand" 	json:"carbrand"  	gorm:"column:carbrand"`
	Carmodel    string 	   `form:"carmodel" 	json:"carmodel"  	gorm:"column:carmodel"`
	Buytime     time.Time  `form:"buytime" 		json:"buytime"  	gorm:"column:buytime"`
}
