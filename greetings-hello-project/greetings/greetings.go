package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	
	if name == "" {
		return "", errors.New("empty name")
	}
	
	

	// In Go, the := operator is a shortcut for declaring and initializng a variable in one line (Go uses the value on the right to determine the variable's type)
	// message := fmt.Sprintf("hi, %v. Welcome!", name)
	// If a name was receiveid, return a greeting that embeds the name
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
