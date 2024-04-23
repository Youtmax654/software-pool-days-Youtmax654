package artists

import (
	rartists "SoftwareGoDay1/artistsBook/repositories/artists"
	vartists "SoftwareGoDay1/artistsBook/views/artists"
)

func DisplayAll() {
	var names []string = rartists.GetAll()
	
	vartists.DisplayAll(names)
}