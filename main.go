package main

import "strings"
func countDistinctWords(messages []string) int {
	distictWords := make(map[string]struct{})

	for _, msg := range messages {
		words := strings.Fields(msg)

		for _, word := range words {
			loweredWord := strings.ToLower(word)
			distictWords[loweredWord] = struct {}{}
		}
	}

	return  len(distictWords)
}
