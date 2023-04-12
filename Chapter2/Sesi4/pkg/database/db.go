package database

import (
	"SWSGolang/app/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	Port     = 5544
	User     = "postgres"
	Password = "Dns215900"
	DBname   = "db_books"
)

func Connection() (*gorm.DB, error) {
	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	Host, Port, User, Password, DBname)

	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
		return nil, err
	}

	db.Debug().AutoMigrate(entity.Book{})
	return db, nil
}