package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := getRandomGreetings(name)
	return message, nil
}

func HelloN(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("Empty names")
	}

	messages := make(map[string]string)

	for _, name := range names {
		message, error := Hello(name)

		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
}

func getRandomGreetings(name string) string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	randomIndex := rand.Intn((len(formats)))
	randomGreetings := formats[randomIndex]

	return fmt.Sprintf(randomGreetings, name)
}
