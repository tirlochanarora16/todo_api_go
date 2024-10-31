package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func EstablishDBConnection() error {
	dbUrl := "host=localhost user=tirlochan password=password dbname=blog_api sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}
