package artists

import (
	"SoftwareGoDay1/artistsBook/models/artist"
	"encoding/json"
	"log"
	"os"
)

func GetAll() []string {
	jsonFile, err := os.ReadFile("artistsBook/data/artists.json")
	if err != nil {
		log.Fatal(err)
	}

	var artists []artist.Artist
	if err := json.Unmarshal(jsonFile, &artists); err != nil {
		log.Fatal(err)
	}

	var names []string
	for _, artist := range artists {
		names = append(names, artist.Name)
	}

	return names
}