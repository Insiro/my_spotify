package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DATABASE_URL"),
	}))
	return db, err
}
