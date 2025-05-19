package main

func main() {
	cards := GetNewDec()
	cards.export("f1")

	cards.shuffle()
	cards.export("f2")
}
