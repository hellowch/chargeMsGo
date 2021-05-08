package models

import (
	"time"
)

type Charger struct {
	Id		 int     `form:"id"          json:"id"         gorm:"column:id;primary_key"`
	Address  string  `form:"address"     json:"address"    gorm:"column:address"`
	Distance int     `form:"distance" 	 json:"distance"   gorm:"column:distance"`
	Usetype  int 	 `form:"usetype" 	 json:"usetype"    gorm:"column:usetype"`
	Details  string  `form:"details" 	 json:"details"    gorm:"column:details"`
}

type ChargerDetails struct {
	Id   		int
	Chargerid   int
	userid		int
 	Details     string
	Time		time.Time
}

type ResChargerDetails struct {
	Id   		int
	Chargerid   int
	userid  	int
	Details     string
	Time		string
}

type ChargerOrder struct {
	Id			int
	Chargerid	int
	Userid		int
	Amount      string
	Time		time.Time
	Length		int
}

type ResChargerOrder struct {
	Id			int
	Chargerid	int
	Userid		int
	Amount      string
	Time		string
	Length		string
}

