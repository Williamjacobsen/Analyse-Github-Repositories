package main

const (
	LOGGING = true
	SAVE_RESULTS_TO_FILE = true
)

func main() {
	fileUrls := getAllFileUrls()
	fileExtensions := getFileExtensions()

	analyseFiles(fileUrls, fileExtensions)
}
