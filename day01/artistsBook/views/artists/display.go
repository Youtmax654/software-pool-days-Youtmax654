package artists

import (
	"SoftwareGoDay1/artistsBook/models/artist"
	"fmt"
)

// prints all the names of the artists from a string array
func DisplayAll(names []artist.Artist) {
	// Check if the list of artists is empty
	if len(names) == 0 {
		return
	}

	// Print the list of artists
	fmt.Println("\nHere's your favorite artists:")
	for i, name := range names {
		fmt.Printf("%d %s", i+1, &name)
	}
}