package artists

import (
	"SoftwareGoDay1/artistsBook/models/artist"
	"encoding/json"
	"log"
	"os"
)

// Add an artist to the list
func Create(name string) {
	// Get all the artists
	jsonFile, err := os.ReadFile("artistsBook/data/artists.json")
	if err != nil {
		log.Fatal(err)
	}

	// Decode the json file into an array of artists
	var artists []artist.Artist
	if err := json.Unmarshal(jsonFile, &artists); err != nil {
		log.Fatal(err)
	}

	// Add the new artist to the array
	artists = append(artists, artist.Artist{Name: name})

	// Encode the array of artists into a json file
	artistsJSON, err := json.Marshal(artists)
	if err != nil {
		log.Fatal(err)
	}

	// Write the json file
	if err := os.WriteFile("artistsBook/data/artists.json", artistsJSON, 0644); err != nil {
		log.Fatal(err)
	}
}