package main

import (
	"fmt"
	"sort"
	"strings"
)

// MakeAnagramMap makes anagram map from given slice of strings
func MakeAnagramMap(words []string) map[string][]string {
	// sorted string -> slice of anagrams
	data := make(map[string][]string)
	setOfWords := make(map[string]interface{})
	for _, word := range words {
		word = strings.ToLower(word)
		_, ok := setOfWords[word]
		if ok {
			continue
		}
		setOfWords[word] = struct{}{}
		stringRunes := []rune(word)
		sort.Slice(stringRunes, func(i int, j int) bool { return stringRunes[i] < stringRunes[j] })
		sortedString := string(stringRunes)
		data[sortedString] = append(data[sortedString], word)
	}
	anagramMap := make(map[string][]string)
	// fill result map
	for _, val := range data {
		anagramMap[val[0]] = val[1:]
	}
	return anagramMap
}

func main() {
	fmt.Printf("%v", MakeAnagramMap([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}))
}
