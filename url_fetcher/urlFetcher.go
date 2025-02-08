package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const OUTPUT_FOLDER = "output"
const FILE_PERMISSION = 0644

func main() {
	urls := os.Args[1:]

	// If there is no urls provided
	if len(urls) == 0 {
		notifyError(errors.New("no url provided"))
	}

	for _, url := range urls {
		body, err := getBodyOfUrl(url)

		if err != nil {
			notifyError(err)
		}

		// Creating the output folder (error ignored)
		os.Mkdir(OUTPUT_FOLDER, FILE_PERMISSION)
		// Creating (updating) the files with response body
		writeError := os.WriteFile(getFileNameFromUrl(url), body, FILE_PERMISSION)

		if writeError != nil {
			notifyError(writeError)
		}
	}
}

func getBodyOfUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func getFileNameFromUrl(url string) string {
	// the prefix is always http:// or https://
	//so we are counting from // + 2 (to skip slashes)
	cutStart := strings.Index(url, "//") + 1
	// The index of last dot (ex: .com)
	cutEndIndex := strings.LastIndex(url, ".")

	return "output/" + url[cutStart+2:cutEndIndex] + ".txt"
}

func notifyError(err error) {
	fmt.Print(err)
	fmt.Fprintf(os.Stderr, "Something went wrong, %v", err)
	os.Exit(1)
}
