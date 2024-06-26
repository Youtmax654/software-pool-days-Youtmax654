package artists

import (
	rartists "SoftwareGoDay1/artistsBook/repositories/artists"
	vartists "SoftwareGoDay1/artistsBook/views/artists"
)

func Create() {
	// Ask the user for the artist's name
	name := vartists.AskName()

	// Call the repository to create the artist
	rartists.Create(name)
}