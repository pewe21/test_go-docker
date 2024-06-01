package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pewe21/go-docker/config"
	"github.com/pewe21/go-docker/handler"
	"github.com/pewe21/go-docker/repository"
	"github.com/pewe21/go-docker/service"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if dbname == "" {
		dbname = "test_go"
	}
	if port == "" {
		port = "5432"
	}
	if host == "" {
		host = "localhost"
	}
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	newDsn := fmt.Sprintf(dsn, host, user, password, dbname, port)
	db, err := config.NewPostgresDB(newDsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	// Cek koneksi
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get raw database connection: %v\n", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	if err := r.Run(":3001"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("Server running on port 3001")
}
