package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	f, err := os.Open(files[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "duplicate err: %v\n", err)
		return
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}

	for text, count := range counts {
		if count > 1 {
			fmt.Printf("The word %v is repeated %d times\n", text, count)
		}
	}

}
