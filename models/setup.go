package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))
	// kondisi jika gagal connect database
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{}) //automigrate ke models

	DB = database
}
