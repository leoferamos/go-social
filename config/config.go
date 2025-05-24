package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	// DatabaseStringConection is the connection string for the database.
	DatabaseStringConection = ""
	// Port is the port on which the application will run.
	Port = 0

	// SecretKey is the secret key used for signing tokens.
	SecretKey []byte
)

// Package config provides configuration settings for the application.
func Init() {
	portStr := os.Getenv("API_PORT")
	if portStr == "" {
		portStr = "5000"
	}
	Port, _ = strconv.Atoi(portStr)

	DatabaseStringConection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
