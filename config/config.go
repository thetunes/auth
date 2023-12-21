package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config initializes the config file
func Config(key string) string {
	// Load environment config
	err := godotenv.Load(".env")

	// Show error that occurs while loading config
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err)
	}
	return os.Getenv(key)
}