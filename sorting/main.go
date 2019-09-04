package main

import (
	"log"
	"sort"
)

func main() {

	fruits := [] string {"banana", "apple", "choco"}
	sort.Sort(byLength(fruits))
	log.Println(fruits)
}

type byLength[] string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}