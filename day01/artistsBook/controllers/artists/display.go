package artists

// import both artists sub-packages from the repositories and views
import (
	rartists "SoftwareGoDay1/artistsBook/repositories/artists"
	vartists "SoftwareGoDay1/artistsBook/views/artists"
)

// controller function that get all the artists names and display their names
func DisplayAll() {
	var names []string = rartists.GetAll()
	
	vartists.DisplayAll(names)
}