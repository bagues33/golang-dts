package database

import (
	"asignment_2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "mac"
	password = ""
	dbname   = "asignment-2"
	db       *gorm.DB
	err      error
)

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Order{}, &models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
