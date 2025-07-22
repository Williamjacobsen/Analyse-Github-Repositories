package main

import (
	"fmt"
	"time"
)

const (
	URL                  = "https://github.com/Williamjacobsen/ClosedAI/tree/main"
	BRANCH               = "main"
	LOGGING              = true
	SAVE_RESULTS_TO_FILE = true
	NUMBER_OF_WORKERS    = 10
)

func main() {

	// Time taken:
	// no worker pool (using recursive depth first search): takes ~16-18 seconds
	// 1 worker: ~16-18 seconds
	// 3 workers: ~6 seconds
	// 10 workers: ~3-4 seconds
	startNow := time.Now()
	fileUrls := discoverAllDirectoriesConcurrently()
	discoverAllFilesTime := time.Since(startNow)

	if LOGGING {
		for _, fileUrl := range fileUrls {
			fmt.Println(fileUrl)
		}
	}

	// Time taken: ~150µs
	startNow = time.Now()
	fileExtensions := getFileExtensions()
	timeToGetFileExtensions := time.Since(startNow)

	// Time taken:
	// Single threaded: ~11 seconds
	// 10 workers: ~1-1.5 seconds
	startNow = time.Now()
	result := analyseFiles(fileUrls, fileExtensions)
	analyseFilesTime := time.Since(startNow)

	// Time taken: ~1.5-2 ms
	startNow = time.Now()
	if SAVE_RESULTS_TO_FILE {
		saveResultToJson(result)
	}
	saveToJsonTime := time.Since(startNow)

	if LOGGING {
		fmt.Println("\nTime taken to discover all files:", discoverAllFilesTime)
		fmt.Println("Time taken to get file extensions:", timeToGetFileExtensions)
		fmt.Println("Time taken to analyse files:", analyseFilesTime)
		fmt.Println("Time taken to save results to json:", saveToJsonTime)
	}
}
