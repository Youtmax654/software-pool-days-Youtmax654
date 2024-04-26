package main

import (
	"SoftwareGoDay2/database"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Get the database URL from the environment
	url := os.Getenv("DB_URL")

	// Initialize the database
	client, err := database.NewDatabase(url)

	if err != nil {
		log.Printf("Failed to initialize database: %v", err)
	}

	if client != nil {
		fmt.Println("Database is ready")
	}
}