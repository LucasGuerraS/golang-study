package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request greeting message
	message, err := greetings.Hello("lucas")
	// If an error is returned, print to console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// If no error is returned print message
	fmt.Println(message)
}
