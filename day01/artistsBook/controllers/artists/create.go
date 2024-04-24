package artists

import (
	rartists "SoftwareGoDay1/artistsBook/repositories/artists"
	vartists "SoftwareGoDay1/artistsBook/views/artists"
	"fmt"
)

func Create() {
	// Ask the user for the artist's name
	name := vartists.AskName()

	// Call the repository to create the artist
	rartists.Create(name)
	
	// Display a message to confirm the artist has been added
	fmt.Println()
	fmt.Println(name, " has been added to your favorite artists!")
}