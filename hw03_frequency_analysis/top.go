package hw03frequencyanalysis

import (
	"math"
	"sort"
	"strings"
)

type pair struct {
	key   string
	value int
}

func Top10(input string) []string {
	words := strings.Fields(input)

	wordsFrequencies := make(map[string]int)

	for _, word := range words {
		if _, ok := wordsFrequencies[word]; ok {
			wordsFrequencies[word]++
		} else {
			wordsFrequencies[word] = 1
		}
	}

	var pairs []pair

	for key, value := range wordsFrequencies {
		pairs = append(pairs, pair{key, value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].value != pairs[j].value {
			return pairs[i].value > pairs[j].value
		} else {
			return pairs[i].key < pairs[j].key
		}
	})

	var output []string

	for i := 0; i < int(math.Min(float64(10), float64(len(pairs)))); i++ {
		output = append(output, pairs[i].key)
	}

	return output
}
