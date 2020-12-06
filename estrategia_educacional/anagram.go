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
	hasAnamagram := func(word, sentence string) bool {
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

		for _, item := range strings.Split(sentence, " ") {
			if isAnamagram(word, item) {
				return true
			}
		}
		return false
	}

	countCombinations := func(words []string, sentence string) int {
		counts := []int{}
		for i := range words {
			if hasAnamagram(words[i], sentence) {
				counts = append(counts, 2)
			}
		}
		count := 1
		for i := range counts {
			count = count * counts[i]
		}
		return count
	}
	output := []int64{}
	for i := range sentences {
		output = append(output, int64(countCombinations(wordSet, sentences[i])))
	}
	return output
}
