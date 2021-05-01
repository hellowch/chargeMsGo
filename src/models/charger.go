package models

import (
	"time"
)

type Charger struct {
	id		 int
	address  string
	distance int
	usetype 	 string
	details  string
}

type ChargerDetails struct {
	id   		int
	chargerid   int
	details     string
	time		time.Time
}

type ChargerOrder struct {
	id			int
	chargerid	int
	userid		int
	amount      string
	time		time.Time
	length		int
}
