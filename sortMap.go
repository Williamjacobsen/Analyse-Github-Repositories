package main

import "fmt"

func sortMap(_map map[string]int) []LanguageLineCount {
	pairs := make([]LanguageLineCount, 0, len(_map))
	for k, v := range _map {
		pairs = append(pairs, LanguageLineCount{Extension: k, Lines: v})
	}

	fmt.Println()

	quickSort(pairs, 0, len(pairs)-1)

	return pairs
}

func quickSort(data []LanguageLineCount, lowIndex int, highIndex int) {
	if lowIndex < highIndex {
		p := partition(data, lowIndex, highIndex)
		quickSort(data, lowIndex, p-1)
		quickSort(data, p+1, highIndex)
	}
}

func partition(data []LanguageLineCount, lowIndex int, highIndex int) int {
	pivot := data[highIndex].Lines
	i := lowIndex
	for j := lowIndex; j <= highIndex; j++ {
		if data[j].Lines > pivot {
			swap(data, i, j)
			i++
		}
	}
	swap(data, i, highIndex)
	return i
}

func swap(data []LanguageLineCount, i int, j int) {
	data[i], data[j] = data[j], data[i]
}
