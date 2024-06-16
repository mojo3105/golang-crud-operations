package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// postgres://bbbvqcim:a2VOaYJgkMySOZF4YI7qzUsGL4gIwnau@snuffleupagus.db.elephantsql.com/bbbvqcim
	dsn := os.Getenv("DB_URL")
	fmt.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
                panic("Failed to connect to database!")
        }

        err = database.AutoMigrate(&Book{})
        if err != nil {
                return
        }

        DB = database
}