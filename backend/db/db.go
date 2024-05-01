package db

import (
	"backend/models"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

var dbInstance *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	lock.Lock()
	defer lock.Unlock()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to create Postgres connection")
	}

	// DB Migration
	err = database.AutoMigrate(&models.AppInfo{}, &models.DNSLog{})
	if err != nil {
		log.Fatalln("Failed to auto-migrate models")
	}

	log.Printf("New PostgresDB Connection Created")
	return database, err
}
