package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Error: not found port")
	}
	return port
}
