package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect() {
	// Replace with your PostgreSQL credentials
	dsn := "host=localhost user=postgres dbname=postgres sslmode=disable"
	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to the database!")
	}
	DB = database
	fmt.Println("Database connected!")
}

func GetDB() *gorm.DB {
	return DB
}
