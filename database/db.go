package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func EstablishDBConnection() error {
	dbUrl := "host=localhost user=tirlochan password=password dbname=todo_app sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}

func RunAutomigration() error {
	err := DB.AutoMigrate(&User{}, &Task{})

	if err != nil {
		log.Fatal("error running the auto migration function!", err)
		return err
	}

	return nil
}
