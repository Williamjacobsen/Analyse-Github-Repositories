package main

import (
	"fmt"
	"slices"
)

const (
	LOGGING = true
)

func main() {
	fileUrls := getAllFileUrls()
	fileExtensions := getFileExtensions()

	fileLines := map[string]int{}

	for _, url := range fileUrls {
		if slices.Contains(fileExtensions, getFileExtensionFromUrl(url)) {
			fileLines[url] = amountOfLinesInFile(url)

			fmt.Println(url)
			fmt.Printf("Amount of lines in file: %d\n\n", fileLines[url])

		}
	}

	fmt.Printf("Total amount of lines written: %d\n", totalAmountOfLines(fileLines))
}
