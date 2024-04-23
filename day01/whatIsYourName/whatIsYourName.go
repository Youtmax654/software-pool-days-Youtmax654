package whatisyourname

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WhatIsYourName() {
	fmt.Println("What is your name ?")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nYour name is ", line)
}