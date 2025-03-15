package database

import (
	"log"

	"github.com/RenuBhati/OnlineCodeEditor/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error

	DB, err = gorm.Open(sqlite.Open("code_editor.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)

	}
	log.Println("Connected to the database successfully")

	err = DB.AutoMigrate(&models.File{})
	if err != nil {
		log.Fatal("Failed to migrate database \n", err)

	}
	log.Println("Successfully Database migration")

}
