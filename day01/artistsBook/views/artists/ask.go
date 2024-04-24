package artists

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AskName() string {
	// Ask the user for the artist's name
	fmt.Println("\nWhat's the artist's name?")
	fmt.Print("\n> ")

	// Read the user input
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	// Check if there is an error
	if err != nil {
		log.Fatal(err)
	}
	
	// Remove white spaces
	name = strings.TrimSpace(name)

	return name
}