package database

import (
	"log"

	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func New() *Database {
	dbURL := "postgres://postgres:mysecretpassword@localhost:5432"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Database{
		DB: db,
	}
}

func (db *Database) Migrate() {
	db.DB.AutoMigrate(&models.Book{}, &models.Author{})
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
