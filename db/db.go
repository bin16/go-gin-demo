package db

import (
	"log"
	"os"

	"github.com/bin16/go-gin-demo/models"
	"github.com/glebarez/sqlite" // pure go, no cgo
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if err := connectToDatabase(); err != nil {
		log.Fatalf("failed to open database: %v\n", err)
	}

	if err := databaseAutoMigrate(); err != nil {
		log.Fatalf("failed to migrate: %v\n", err)
	}
}

func connectToDatabase() error {
	dbType := os.Getenv("DB")
	dbURL := os.Getenv("DB_URL")
	if dbType == "postgres" {
		dbURL := os.Getenv(dbURL)
		db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			return err
		}
		DB = db

		return nil
	}

	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func databaseAutoMigrate() error {
	return DB.AutoMigrate(
		&models.Profile{},
		&models.User{},
		&models.Note{},
	)
}
