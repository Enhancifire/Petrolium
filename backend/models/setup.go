package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	DBURL := ""
	DB, err := gorm.Open("", DBURL)

	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&User{})
}
