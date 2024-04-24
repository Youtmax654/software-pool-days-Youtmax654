package artistsbook

import (
	"SoftwareGoDay1/artistsBook/controllers/artists"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Router() {
	fmt.Println("Welcome into your Artists Book!")
	for {
		// List input options
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1 - List my favorite artists")
		fmt.Println("2 - Add an artist to my favorite")
		fmt.Println("3 - Update an artist in my favorite")
		fmt.Println("4 - Delete an artist from my favorite")
		fmt.Println("5 - Leave")

		// Create a loop to keep the program running
		for {
			fmt.Print("\n> ")
			// Read the user input
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')

			// Check if there is an error
			if err != nil {
				log.Fatal(err)
			}

			// Remove white spaces
			line = strings.TrimSpace(line)

			// Check the user input
			if line == "1" {
				artists.DisplayAll()
				break
			} else if line == "2" {
				artists.Create()
				break
			} else if line == "3" {
				artists.Update()
				break
			} else if line == "4" {
				artists.Delete()
				break
			} else if line == "5" {
				fmt.Println("\nSee you!")
				// Exit the loop
				return
			} else {
				fmt.Println("\nTip 1, 2, 3, 4 or 5.")
			}
		}
	}
}
