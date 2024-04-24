package artists

import (
	rartists "SoftwareGoDay1/artistsBook/repositories/artists"
	vartists "SoftwareGoDay1/artistsBook/views/artists"
)

func Delete() {
	// Ask the user for the artist's name
	name := vartists.AskRemoveName()

	// Delete the artist
	rartists.Delete(name)
}