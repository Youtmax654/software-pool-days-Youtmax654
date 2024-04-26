package main

import (
	"SoftwareGoDay2/database"
	"SoftwareGoDay2/ent"
	"context"
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

	ctx := context.Background()

	if client != nil {
		artist, _ := client.CreateArtist(ctx, "Florida_Test", "")
		log.Printf("Artist created: %v", artist)

		updated, _ := client.UpdateArtist(ctx, &ent.Artist{ID: artist.ID, Name: "Florida_Test", Nationality: "us"})
		log.Printf("Artist updated: %v", updated)
	}
}
