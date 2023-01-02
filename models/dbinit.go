package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to connect .env")
	}
	db_dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(db_dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
	db.AutoMigrate(
		&Books{},
		&Users{},
	)
	log.Println("Migrating DB Successfully")
}
