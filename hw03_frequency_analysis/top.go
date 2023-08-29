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
		_, ok := wordsFrequencies[word]
		if ok {
			wordsFrequencies[word]++
		} else {
			wordsFrequencies[word] = 1
		}
	}

	pairs := make([]pair, len(wordsFrequencies))

	i := 0
	for key, value := range wordsFrequencies {
		pairs[i] = pair{key, value}
		i++
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
