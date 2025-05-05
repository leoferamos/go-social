package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseStringConection is the connection string for the database.
	DatabaseStringConection = ""
	// Port is the port on which the application will run.
	Port = 0
)

// Package config provides configuration settings for the application.
func Init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000 // Default port
	}
	DatabaseStringConection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
