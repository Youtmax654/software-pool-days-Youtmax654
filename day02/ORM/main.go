package main

import (
	"SoftwareGoDay2/database"
	"SoftwareGoDay2/ent"
	"context"
	"log"
	"os"

	"github.com/google/uuid"
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

	ctx := context.Background()

	var contact *ent.Contact = &ent.Contact{
		ID: uuid.New(),
		Email: "youtmax654@email.com",
		Phone: "1234567890",
	}

	update, err := client.UpdateContact(ctx, contact)
	if err != nil {
		log.Printf("Failed to update contact: %v", err)
	}

	log.Printf("Updated contact: %v", update)
}
