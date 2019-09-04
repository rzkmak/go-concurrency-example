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

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }