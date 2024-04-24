package artists

import (
	rartists "SoftwareGoDay1/artistsBook/repositories/artists"
	vartists "SoftwareGoDay1/artistsBook/views/artists"
)

func Update() {
	// Ask the user for the artist's name and the new name
	name, newName := vartists.AskNewName()

	// Update the artist
	rartists.Update(name, newName)
}