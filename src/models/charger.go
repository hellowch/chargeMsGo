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
	Bdlng	 string  `form:"bdlng" 		 json:"bdlng"      gorm:"column:bdlng"`
	Bdlab	 string  `form:"bdlab" 	 	 json:"bdlab"      gorm:"column:bdlab"`
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
	Id			int			`form:"id"`
	Chargerid	int			`form:"chargerid"`
	Userid		int			`form:"userid"`
	Amount      string		`form:"amount"`
	Time		time.Time	`form:"time"`
	Length		int			`form:"length"`
}

type ResChargerOrder struct {
	Id			int
	Chargerid	int
	Userid		int
	Amount      string
	Time		string
	Length		string
}

