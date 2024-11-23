package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Babken", "Narek", "Suren"}
	message, error := greetings.HelloN(names)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
