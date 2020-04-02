package gotools

import "strings"

// WordsToMap counts the occurrences of words and returns a map of them
func WordsToMap(s string) map[string]int {
	words := strings.Fields(s)
	wordmap := map[string]int{}
	for _, word := range words {
		wordmap[word]++
	}
	return wordmap
}
