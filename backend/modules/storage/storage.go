package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	// DBURL := ""
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// DB.AutoMigrate(&User{})
}

func Migrate(a any) {
	if err := DB.AutoMigrate(&a); err != nil {
		log.Fatal(err)
	}
}
