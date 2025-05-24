package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HashKey  []byte
	BlockKey []byte
)

func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
