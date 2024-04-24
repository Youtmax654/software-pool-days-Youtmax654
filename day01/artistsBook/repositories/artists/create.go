package artists

import (
	"SoftwareGoDay1/artistsBook/models/artist"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aidarkhanov/nanoid"
)

// Add an artist to the list
func Create(name string) {
	// Get all the artists
	jsonFile, err := os.ReadFile("artistsBook/data/artists.json")
	if err != nil {
		log.Fatal(err)
	}

	var artists []artist.Artist
	
	// Check if the json file is not empty
	if (len(jsonFile) > 0) {
		// Decode the json file into an array of artists
		if err := json.Unmarshal(jsonFile, &artists); err != nil {
			log.Fatal(err)
		}

		for _, artist := range artists {
			if (artist.Name == name) {
				fmt.Println()
				fmt.Println(name, "is already in your favorite artists!")
				return
			}
		}
	}
	
	// Generate a unique id for the new artist
	var id string = nanoid.New()
	// Add the new artist to the array
	artists = append(artists, artist.Artist{Id: id, Name: name})

	// Encode the array of artists into a json file
	artistsJSON, err := json.Marshal(artists)
	if err != nil {
		log.Fatal(err)
	}

	// Write the json file
	if err := os.WriteFile("artistsBook/data/artists.json", artistsJSON, 0644); err != nil {
		log.Fatal(err)
	}

	// Display a message to confirm the artist has been added
	fmt.Println()
	fmt.Println(name, "has been added to your favorite artists!")
}