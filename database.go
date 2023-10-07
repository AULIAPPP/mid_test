package database

import (
	"fmt"
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "username:password@tcp(localhost:3306)/database_name?parseTime=true"
	// Ganti username, password, dan database_name sesuai dengan pengaturan database Anda.
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{})
}
