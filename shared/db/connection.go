package db

import (
	"fmt"
	"log"

	"github.com/workloom/shared/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func Init() {
    dsn := "host=localhost user=postgres password=password dbname=workloom port=5432 sslmode=disable TimeZone=Asia/Kolkata"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&models.User{})
	
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    } else {
		fmt.Println("DataBase Connected Successfully")
	}
}
