package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(isPangram("input "))
}

func isPangram(word string) string {
	var filteredString string
	filteredString = strings.ReplaceAll(strings.ToLower(word), " ", "")
	charArray := []rune (filteredString)
	mapCharArray := make(map[int]bool)
	for _, w:= range charArray {
		mapCharArray[int(w)] = true
	}
	for i:= 97; i<=122; i++ {
		if !mapCharArray[i] {
			return "not pangram"
		}
	}
	return "pangram"
}
