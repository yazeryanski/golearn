package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const EXPORT_DIR string = "export"
const DEFAULT_PERMISSION os.FileMode = 0666
const SEPARATOR = "\n"

var suits []string = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
var values []string = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

type Deck []string

// Constructors
func GetNewDec() Deck {
	var deck Deck

	for _, suit := range suits {
		for _, val := range values {
			deck = append(deck, val+" of "+suit)
		}
	}

	return deck
}

func GetNewDeckFromFile(filePath string) (Deck, error) {
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	content := string(bytes)
	result := strings.Split(content, SEPARATOR)

	return result, nil
}

// Methods
func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i+1, ")", card)
	}
}

func (d *Deck) deal(amount uint8) Deck {
	cardsToDeal := (*d)[:amount]
	*d = (*d)[amount:]

	return cardsToDeal
}

func (d Deck) export(fileName string) error {
	os.Mkdir(EXPORT_DIR, os.ModePerm)

	filePath := path.Join(EXPORT_DIR, fileName)
	body := []byte(strings.Join(d, "\n"))

	return os.WriteFile(filePath, body, DEFAULT_PERMISSION)
}

func (d *Deck) shuffle() {
	*d = ShuffleArray(*d)
}
