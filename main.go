package main

import (
	"fmt"
)

const (
	LOGGING = true
)

func main() {
	fileUrls := getAllFileUrls()

	for _, fileUrl := range fileUrls {
		fmt.Println(fileUrl)
	}
}
