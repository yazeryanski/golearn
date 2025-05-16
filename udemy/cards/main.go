package main

import "fmt"

func main() {
	cards := GetNewDec()

	tony := cards.deal(3)
	john := cards.deal(3)

	tony.print()
	john.print()

	// #region export
	err := cards.export("test.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("File successfully exported")
	// #endregion

	// #region import
	newDeck, err := GetNewDeckFromFile("export/test.txt")

	if err != nil {
		panic(err)
	}

	bob := newDeck.deal(5)

	fmt.Println(bob)
	// #endregion
}
