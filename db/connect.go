package db

import "github.com/jinzhu/gorm"

var dbCon *gorm.DB

func init() {
	dsn := "root:mSsDb*4297pC@tcp(9.134.177.71:3306)/bull_mock?charset=utf8mb4&parseTime=True&loc=Local"
	dbCon, _ = gorm.Open("mysql", dsn)
}

func dstory() {

}
