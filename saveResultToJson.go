package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func saveResultToJson(result AnalysisResult) {
	encoder, file := createJsonOutputFile()
	defer file.Close()

	if err := encoder.Encode(result); err != nil {
		log.Fatalf("could not encode to json file: %s", err)
	}

	if LOGGING {
		fmt.Println("Saved results to analysis_result.json")
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
