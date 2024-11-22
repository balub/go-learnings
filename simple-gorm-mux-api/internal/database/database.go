package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func New() *Database {
	db, err := gorm.Open(sqlite.Open(os.Getenv("FILENAME")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Database{
		DB: db,
	}
}

func (db *Database) Migrate() {
	db.DB.AutoMigrate()
}

func (db *Database) Health() map[string]string {
	stats := make(map[string]string)
	stats["health"] = "good"

	return stats
}

func (db *Database) Close() error {
	log.Printf("Disconnected from database")
	return nil
}

func (db *Database) GetDB() *gorm.DB {
	return db.DB
}
