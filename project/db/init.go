package db

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func Connect() error{
	var dbURL string = "postgres://postgres:postgres@localhost:5432"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect database")
	}
	DbConn = db
	return nil
}

