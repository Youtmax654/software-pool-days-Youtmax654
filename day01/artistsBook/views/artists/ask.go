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

func AskNewName() (string, string) {
	// Ask the user for the artist's name to update
	fmt.Println("\nWhat's the name of the artist you want to update?")
	fmt.Print("\n> ")

	// Read the user input for the artist's name
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	// Check if there is an error
	if err != nil {
		log.Fatal(err)
	}

	// Remove white spaces
	name = strings.TrimSpace(name)

	// Ask the user for the new name
	fmt.Print("\nWhat will be the new name of ", name, "?")
	fmt.Println()
	fmt.Print("\n> ")

	// Read the user input for the new name
	newName, err := reader.ReadString('\n')

	// Check if there is an error
	if err != nil {
		log.Fatal(err)
	}

	// Remove white spaces
	newName = strings.TrimSpace(newName)

	return name, newName
}

func AskRemoveName() string {
	// Ask the user for the artist's name to remove
	fmt.Println("\nWhat's the name of the artist you want to remove?")
	fmt.Print("\n> ")

	// Read the user input for the artist's name
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