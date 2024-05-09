package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting from external function
	message := greetings.Hello("Lucas")
	fmt.Println(message)
}
