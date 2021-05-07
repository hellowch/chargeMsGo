package models

type Community struct {
	Id			int
	Userid		int
	Details		string
	Detailsid   int
}

type ResCommunity struct {
	Id			int
	Userid		int
	Details		string
	Detailsid   int
	Name        string
	Avatar      string
}
