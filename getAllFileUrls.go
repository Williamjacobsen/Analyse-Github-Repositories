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
			endIndex = i + 1
			break
		}
	}

	json := html[startIndex:endIndex]
	return json
}

func getDirectories(items gjson.Result, baseUrl string, rawFileUrl string) ([]string, []string) {
	var fileUrls []string
	var directoryUrls []string

	items.ForEach(func(_, item gjson.Result) bool {
		path := url.PathEscape(item.Get("path").Str)
		contentType := item.Get("contentType").Str

		switch contentType {
		case "file":
			fileUrls = append(fileUrls, rawFileUrl+"/"+path)
		case "directory":
			directoryUrls = append(directoryUrls, baseUrl+"/"+path)
		}

		return true
	})

	return directoryUrls, fileUrls
}

func getDirectoriesWrapper(url string, rawFileUrl string, baseUrl string) ([]string, []string) {
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
	also append found files to allFileUrls
*/
func recursiveDirectoryDepthFirstSearch(directories []string, rawFileUrl string, baseUrl string, allFileUrls *[]string) {
	if len(directories) == 0 {
		return
	}

	for _, directory := range directories {
		if LOGGING {
			fmt.Println("Visiting:", directory)
		}

		subDirectories, fileUrls := getDirectoriesWrapper(directory, rawFileUrl, baseUrl)

		*allFileUrls = append(*allFileUrls, fileUrls...)

		recursiveDirectoryDepthFirstSearch(subDirectories, rawFileUrl, baseUrl, allFileUrls)
	}
}

func getAllFileUrls() []string {
	/*
		At the root directory the json that should be search for is `{"props":{"initialPayload":`
		And the JSON path to the data is "props.initialPayload.tree.items".

		At a nested directory it should be `{"payload":{`
		And the JSON path is "payload.tree.items"
	*/

	// TODO: Error if it doesn't have /tree/main surffix

	repo := strings.TrimPrefix(URL, "https://github.com/")
	repo = strings.TrimSuffix(repo, "/tree/main")

	rawFileUrl := "https://raw.githubusercontent.com/" + repo + "/refs/heads/" + BRANCH

	html := getHtml(URL)
	json := getJson(html, `{"props":{"initialPayload":`)
	items := gjson.Get(json, "props.initialPayload.tree.items")

	rootDirectoriesUrls, rootFileUrls := getDirectories(items, URL, rawFileUrl)

	var allFileUrls []string
	allFileUrls = append(allFileUrls, rootFileUrls...)

	recursiveDirectoryDepthFirstSearch(rootDirectoriesUrls, rawFileUrl, URL, &allFileUrls)

	return allFileUrls
}
