package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
var (
	db *gorm.DB
)
func StartDatabase() {
	db, err := gorm.Open("mysql", "root: @(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.AutoMigrate(&User{})

}
