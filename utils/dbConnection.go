package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB Global variable
var DB *gorm.DB

// To initilize the postgres db connection
func Init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=goMux_db port=5430 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	log.Println("Connect to the DB")
	// DB.AutoMigrate(&models.User{})
}
