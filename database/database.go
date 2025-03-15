package database

import (
	"log"
	"os"

	"github.com/RenuBhati/OnlineCodeEditor/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {

	var err error

	DB, err = gorm.Open(sqlite.Open("online_code_editor.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")

	err = DB.AutoMigrate(&models.File{}, &models.FileShare{}, &models.GitCommit{})
	if err != nil {
		log.Fatal("Failed to migrate database \n", err.Error())
		os.Exit(2)
	}
	log.Println("Successfully Database migration")

}
