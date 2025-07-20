package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getFileExtensions() []string {
	file, err := os.Open("fileExtensions.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileExtensions []string

	for scanner.Scan() {
		line := scanner.Text()
		fileExtensions = append(fileExtensions, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return fileExtensions
}

func getFileExtensionFromUrl(url string) string {
	fileExtensionStart := strings.LastIndex(url, ".")
	if fileExtensionStart == -1 {
		log.Fatalf("no file extension for file: %s", url)
	}

	return url[fileExtensionStart:]
}
