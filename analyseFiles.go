package main

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

type AnalysisResult struct {
	FileLines            map[string]int `json:"files and their lines"`
	LineCountPerLanguage map[string]int `json:"lines per language"`
	TotalLineCount       int            `json:"total amount of lines"`
}

func analyseFiles(fileUrls []string, fileExtensions []string) AnalysisResult {
	fileLines := map[string]int{}

	type job struct {
		url string
	}

	type result struct {
		url       string
		lineCount int
	}

	jobs := make(chan job, len(fileUrls))
	results := make(chan result, len(fileUrls))

	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for j := range jobs {
			if slices.Contains(fileExtensions, getFileExtensionFromUrl(j.url)) {
				lineCount := amountOfLinesInFile(j.url)

				if LOGGING {
					fmt.Println(j.url)
					fmt.Printf("Amount of lines in file: %d\n\n", lineCount)
				}

				results <- result{url: j.url, lineCount: lineCount}
			}
		}
	}

	wg.Add(NUMBER_OF_WORKERS)
	for w := 1; w <= NUMBER_OF_WORKERS; w++ {
		go worker()
	}

	go func() {
		for _, url := range fileUrls {
			jobs <- job{url: url}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fileLines[result.url] = result.lineCount
	}

	lineCountPerLanguage := lineCountPerLanguage(fileLines)
	totalLineCount := totalAmountOfLines(fileLines)

	return AnalysisResult{
		FileLines:            fileLines,
		LineCountPerLanguage: lineCountPerLanguage,
		TotalLineCount:       totalLineCount,
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

func lineCountPerLanguage(fileLines map[string]int) map[string]int {
	linesPerLang := map[string]int{}

	for url, lines := range fileLines {
		linesPerLang[getFileExtensionFromUrl(url)] += lines
	}

	return linesPerLang
}
