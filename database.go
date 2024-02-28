package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func databaseSetup(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in New | Establishing a connection", err)
	}

	err = db.AutoMigrate(&Gift{})
	if err != nil {
		fmt.Println("Couldn't Automigrate Database.", err)
	}
	Database = db
}
