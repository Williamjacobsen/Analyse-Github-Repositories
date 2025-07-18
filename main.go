package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

func getHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

// Assumes the JSON object starts with '{' / it is not an array
func getJson(html string, startOfJson string) string {
	startIndex := strings.Index(html, startOfJson)
	endIndex := -1

	currentlyOpenObjects := 0
	for i := startIndex; i < len(html); i++ {
		switch html[i] {
			case '{':
				currentlyOpenObjects++
			case '}':
				currentlyOpenObjects--
		}

		if currentlyOpenObjects == 0 {
			endIndex = i+1
			break
		}
	}

	json := html[startIndex:endIndex]
	return json
}

func getDirectories(items gjson.Result, baseUrl string, rawFileUrl string) []string {
	var fileUrls []string
	var directoryUrls []string

	items.ForEach(func(_, item gjson.Result) bool {
		
		//name := item.Get("name").Str
		path := url.PathEscape(item.Get("path").Str)
		contentType := item.Get("contentType").Str
		
		//fmt.Printf("Name: %s, Path: %s, contentType: %s\n", name, path, contentType)
		
		switch contentType {
			case "file":
				fileUrls = append(fileUrls, rawFileUrl + "/" + path)
			case "directory":
				directoryUrls = append(directoryUrls, baseUrl + "/" + path)
		}

		return true
	})

	//for _, fileUrl := range fileUrls {
	//	fmt.Println("File URL: " + fileUrl)
	//}

	//for _, directoryUrl := range directoryUrls {
	//	fmt.Println("directory URL: " + directoryUrl)
	//}

	return directoryUrls
}

func getDirectoriesWrapper(url string, rawFileUrl string, baseUrl string) []string {
	html := getHtml(url)

	json := getJson(html, `{"payload":{`)

	items := gjson.Get(json, "payload.tree.items")

	return getDirectories(items, baseUrl, rawFileUrl)
}

/*
given an array of directories (from root)

if no directories return
otherwise find all sub directories of each directory - for each:
	call self with array of newly discovered sub directories

*/
func recursiveDirectoryDepthFirstSearch(directories []string, rawFileUrl string, baseUrl string) {
	if len(directories) == 0 {
		return
	}

	for _, directory := range directories {
		fmt.Println("Visiting:", directory)

		subDirectories := getDirectoriesWrapper(directory, rawFileUrl, baseUrl)
		recursiveDirectoryDepthFirstSearch(subDirectories, rawFileUrl, baseUrl)
	}
}

func main() {

	/*
	At the root directory the json that should be search for is `{"props":{"initialPayload":`
	And the JSON path to the data is "props.initialPayload.tree.items".

	At a nested directory it should be `{"payload":{`
	And the JSON path is "payload.tree.items"
	*/

	url := "https://github.com/Williamjacobsen/ClosedAI/tree/main"
	branch := "main"

	repo := strings.TrimPrefix(url, "https://github.com/")
	repo = strings.TrimSuffix(repo, "/tree/main")

	rawFileUrl := "https://raw.githubusercontent.com/" + repo + "/refs/heads/" + branch

	html := getHtml(url)
	json := getJson(html, `{"props":{"initialPayload":`)
	items := gjson.Get(json, "props.initialPayload.tree.items")

	rootDirectoriesUrls := getDirectories(items, url, rawFileUrl)

	recursiveDirectoryDepthFirstSearch(rootDirectoriesUrls, rawFileUrl, url)
}
