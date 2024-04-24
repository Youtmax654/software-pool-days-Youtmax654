package artists

import (
	"SoftwareGoDay1/artistsBook/models/artist"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Delete(name string) {
	jsonFile, err := os.ReadFile("artistsBook/data/artists.json")
	if err != nil {
		log.Fatal(err)
	}

	var artists []artist.Artist
	// Decode the json file into an array of artists
	if err := json.Unmarshal(jsonFile, &artists); err != nil {
		log.Fatal(err)
	}

	var found bool
	// Delete the artist
	for i, artist := range artists {
		if artist.Name == name {
			// Append the artists before and after the artist to delete
			artists = append(artists[:i], artists[i+1:]...)
			found = true
		}
	}

	// Check if the artist is in the favorite list
	if !found {
		fmt.Println()
		fmt.Println(name, "is not in your favorite list.")
		return
	}

	// Encode the array of artists into a json file
	artistsJSON, err := json.Marshal(artists)
	if err != nil {
		log.Fatal(err)
	}

	// Write the json file
	if err := os.WriteFile("artistsBook/data/artists.json", artistsJSON, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(name, "has been successfully deleted.")
}