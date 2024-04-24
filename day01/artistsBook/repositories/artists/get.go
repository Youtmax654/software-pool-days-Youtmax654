package artists

import (
	"SoftwareGoDay1/artistsBook/models/artist"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// get all the artists and return their names
func GetAll() []artist.Artist {
	// read the json file that contains the artists
	jsonFile, err := os.ReadFile("artistsBook/data/artists.json")
	if err != nil {
		log.Fatal(err)
	}

	// Check if the json is empty
	if len(jsonFile) == 0 {
		fmt.Println("\nYour list of favorite artists is empty. Add an artist to get started!")
		return nil
	}

	var artists []artist.Artist
	// decode the json file into an array of artists
	if err := json.Unmarshal(jsonFile, &artists); err != nil {
		log.Fatal(err)
	}
	
	// Check if the list is empty
	if len(artists) == 0 {
		fmt.Println("\nYour list of favorite artists is empty. Add an artist to get started!")
		return nil
	}

	// var names []string
	// // get the names of the artists
	// for _, artist := range artists {
	// 	names = append(names, artist.Name)
	// }

	return artists
}