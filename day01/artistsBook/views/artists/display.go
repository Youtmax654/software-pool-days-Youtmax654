package artists

import (
	"fmt"
)

// prints all the names of the artists from a string array
func DisplayAll(names []string) {
	// Check if the list of artists is empty
	if len(names) == 0 {
		return
	}

	// Print the list of artists
	fmt.Println("\nHere's your favorite artists:")
	for i := 1; i <= len(names); i++ {
		fmt.Println("--", i, "--", names[i-1])
	}
}