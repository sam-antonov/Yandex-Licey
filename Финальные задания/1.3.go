package main

import (
	"fmt"
	"strings"
	"maps"
)

func getTopWords(wordMap map[string]int, n int) []string {
	var topWords []string
	var wordMapCopy map[string]int
	maps.Copy(wordMap, wordMapCopy)
	for _ = range n {
		max := 0
		var maxWord string
		for key, _ := range wordMap {
			if wordMap[key] > max {
				max = wordMap[key]
				maxWord = key
			}
		}
		topWords = append(topWords, maxWord)
		delete(wordMap, maxWord)
	}
	return topWords
}

func AnalyzeText(text string) {
	wordMap := make(map[string]int)
	for _, symb := range ".,!?" {
		text = strings.ReplaceAll(text, string(symb), "")
	}
	textSlice := strings.Split(text, " ")

	for i := 0; i < len(textSlice); i++ {
		wordMap[strings.ToLower(textSlice[i])] += 1
	}


	wordMapCopy := make(map[string]int)

	maps.Copy(wordMapCopy, wordMap)

	topWords := getTopWords(wordMapCopy, 5)

	fmt.Println("Количество слов:", len(textSlice))
	fmt.Println("Количество уникальных слов:", len(wordMap))
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", topWords[0], wordMap[topWords[0]])
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for i := 0; i < len(topWords); i++ {
		fmt.Printf("\"%s\": %d раз\n", topWords[i], wordMap[topWords[i]])
	}
}