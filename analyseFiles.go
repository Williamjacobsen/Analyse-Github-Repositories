package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type AnalysisResult struct {
	FileLines      map[string]int `json:"files and their lines"`
	TotalLineCount int            `json:"total amount of lines"`
}

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

	var totalLineCount int

	if LOGGING && SAVE_RESULTS_TO_FILE {
		totalLineCount = totalAmountOfLines(fileLines)
	}

	if LOGGING {
		fmt.Printf("Total amount of lines: %d\n", totalLineCount)
	}

	if SAVE_RESULTS_TO_FILE {
		result := AnalysisResult{
			FileLines:      fileLines,
			TotalLineCount: totalLineCount,
		}

		encoder, file := createJsonOutputFile()
		defer file.Close()

		if err := encoder.Encode(result); err != nil {
			log.Fatalf("could not encode to json file: %s", err)
		}
	}
}

func createJsonOutputFile() (*json.Encoder, *os.File) {
	file, err := os.Create("analysis_result.json")
	if err != nil {
		log.Fatalf("could not create analysis_result.json: %s", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder, file
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
