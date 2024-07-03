package types

import "time"

// UserVisit 访客记录
type UserVisit struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	URL            string    `json:"url"`
	Referrer       string    `json:"referrer"`
	ClientID       string    `json:"clientID"`
	UserID         uint      `json:"userID"`
	VisitTime      time.Time `json:"visitTime"`
	IP             string    `json:"ip"`
	DeviceWidth    int       `json:"deviceWidth"`
	DeviceHeight   int       `json:"deviceHeight"`
	BrowserName    string    `json:"browserName"`
	BrowserVersion string    `json:"browserVersion"`
	DeviceModel    string    `json:"deviceModel"`
	Country        string    `json:"country"`
	Language       string    `json:"language"`
	OSName         string    `json:"osName"`
	OSVersion      string    `json:"osVersion"`
}

// PVPerDay 每天的pv
type PVPerDay []struct {
	PV   int    `json:"pv"`
	Date string `gorm:"column:date" json:"date"`
}
