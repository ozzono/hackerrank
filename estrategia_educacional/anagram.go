package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Printf("%+v\n", countSentences([]string{"the", "bats", "tabs", "in", "cat", "act"}, []string{"cat the bats", "in the act", "act tabs in"}))
}

func countSentences(wordSet []string, sentences []string) []int64 {
	countAnamagram := func(word, sentence string) int {
		isAnamagram := func(w1, w2 string) bool {
			if len(w1) != len(w2) || w1 == w2 {
				return false
			}
			mapw1, mapw2 := map[string]string{}, map[string]string{}
			for i := range w1 {
				mapw1[string(w1[i])] = string(w1[i])
				mapw2[string(w2[i])] = string(w2[i])
			}
			out := reflect.DeepEqual(mapw1, mapw2)
			return out
		}

		count := 0
		for _, item := range strings.Split(sentence, " ") {
			if isAnamagram(word, item) {
				count++
			}
		}
		return count
	}

	countCombinations := func(words []string, sentence string) int {
		count := 0
		for i := range words {
			anagrams := countAnamagram(words[i], sentence)
			if anagrams > 0 {
				fmt.Println(anagrams)
				if count == 0 {
					count = 1
				}
				count *= 2 * *&anagrams
				fmt.Printf("count: %d\n", count)
			}
		}
		return count
	}
	output := []int64{}
	for i := range sentences {
		output = append(output, int64(countCombinations(wordSet, sentences[i])))
	}
	return output
}
