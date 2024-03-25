package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// var db *gorm.DB

// func StartDB() {
// 	var err error
// 	db, err = gorm.Open("sqlite3", "final_project.db")
// 	if err != nil {
// 		fmt.Println("Failed to connect to the database:", err)
// 		return
// 	}

// 	db.AutoMigrate(&models.User{})
// 	db.AutoMigrate(&models.Photo{})
// 	db.AutoMigrate(&models.Comment{})
// 	db.AutoMigrate(&models.SocialMedia{})
// }

// func GetDB() *gorm.DB {
// 	return db
// }

type Database struct {
	DB *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{DB: db}
}
