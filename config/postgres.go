package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	Password string
}

// NewPostgresDB creates a new connection to a Postgres database
func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to migrate user: %v\n", err)
	}
	return db, nil
}
