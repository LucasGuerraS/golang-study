package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//If no name was given return with error
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of some greetings
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v Welcome!",
		"Great to see you %v",
		"Hail, %v Well met!",
	}

	// Return a randomly selected randomFormat
	return formats[rand.Intn(len(formats))]
}
