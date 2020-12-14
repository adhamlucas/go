package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	
	if name == "" {
		return "", errors.New("empty name")
	}
	

	// In Go, the := operator is a shortcut for declaring and initializng a variable in one line (Go uses the value on the right to determine the variable's type)
	// message := fmt.Sprintf("hi, %v. Welcome!", name)
	// If a name was receiveid, return a greeting that embeds the name
	message := fmt.Sprintf(randomformat(), name)
	

	return message, nil
}

// init sets initial value sof variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
}

// randoFormat return one fo a set of greeting messages. The returned
// message is selected at random
func randomformat() string {
	formats := []string {
		"hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}