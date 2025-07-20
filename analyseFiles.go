package main

import (
	"fmt"
	"slices"
	"strings"
)

func analyseFiles(fileUrls []string, fileExtensions []string) {
	fileLines := map[string]int{}

	for _, url := range fileUrls {
		if slices.Contains(fileExtensions, getFileExtensionFromUrl(url)) {
			fileLines[url] = amountOfLinesInFile(url)

			if LOGGING {
				fmt.Println(url)
				fmt.Printf("Amount of lines in file: %d\n\n", fileLines[url])
			}
		}
	}

	if LOGGING {
		fmt.Printf("Total amount of lines written: %d\n", totalAmountOfLines(fileLines))
	}
}

func amountOfLinesInFile(url string) int {
	code := getHtml(url)
	lines := strings.Split(code, "\n")
	return len(lines)
}

func totalAmountOfLines(fileLines map[string]int) int {
	totalLines := 0

	for _, linesInFile := range fileLines {
		totalLines += linesInFile
	}
	
	return totalLines
}
