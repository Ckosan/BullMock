package db

import "time"

type MockInfoModel struct {
	ID              uint `gorm:"primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Method          string `gorm:"size:8"`
	Url             string `gorm:"size:256"`
	Request         interface{}
	Response        interface{}
	ReqContentType  string
	RespContentType string
	ForwardUrl      string
	MockRule        string
}
