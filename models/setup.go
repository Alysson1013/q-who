package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


var DB *gorm.DB

func connectDatabe () {
	dsn := "host=localhost user=docker dbname=qwho port=5432 password=12345678"	

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
        panic("Failed to connect to database!")
    }

	database.AutoMigrate(&User{}) 

	DB = database
}