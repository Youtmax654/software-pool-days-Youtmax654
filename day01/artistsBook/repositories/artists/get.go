package artists

import (
	"SoftwareGoDay1/artistsBook/models/artist"
	"encoding/json"
	"log"
	"os"
)

// get all the artists and return their names
func GetAll() []string {
	// read the json file that contains the artists
	jsonFile, err := os.ReadFile("artistsBook/data/artists.json")
	if err != nil {
		log.Fatal(err)
	}

	var artists []artist.Artist
	// decode the json file into an array of artists
	if err := json.Unmarshal(jsonFile, &artists); err != nil {
		log.Fatal(err)
	}

	var names []string
	// get the names of the artists
	for _, artist := range artists {
		names = append(names, artist.Name)
	}

	return names
}