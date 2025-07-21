package main

import (
	"fmt"
	"slices"
	"strings"
)

type AnalysisResult struct {
	FileLines      map[string]int `json:"files and their lines"`
	TotalLineCount int            `json:"total amount of lines"`
}

func analyseFiles(fileUrls []string, fileExtensions []string) AnalysisResult {
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

	totalLineCount := totalAmountOfLines(fileLines)

	result := AnalysisResult{
		FileLines:      fileLines,
		TotalLineCount: totalLineCount,
	}

	return result 
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
