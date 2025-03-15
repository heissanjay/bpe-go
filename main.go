package main

import (
	"fmt"
	"strings"
)

const special_token string = "_"

var vocabulary map[string]int = map[string]int{}

func splitTextIntoWords(corpus string) []string {
	return strings.Split(corpus, " ")
}

func getInitialTokens(words []string) [][]string {
	tokens := make([][]string, 0, len(words))

	for _, word := range words {
		word += special_token
		chars := make([]string, 0, len(word))

		for j := 0; j < len(word); j++ {
			chars = append(chars, string(word[j]))
		}

		tokens = append(tokens, chars)
	}
	return tokens
}

func countPairs(tokens [][]string) map[string]int {
	pairs := make(map[string]int)

	for _, word := range tokens {
		for i := 0; i < len(word)-1; i++ {
			pair := string(word[i]) + string(word[i+1])
			pairs[pair]++
		}
	}
	return pairs
}

func findMostFreqPairs(freqMap map[string]int) (max_freq int, best_pair string) {
	max_freq = 0
	for pair, freq := range freqMap {
		if freq > max_freq {
			max_freq = freq
			best_pair = pair
		}
	}

	return max_freq, best_pair
}

func updateVocabulary(tokens [][]string) {
	vocabulary = make(map[string]int)
	for _, token := range tokens {
		for _, t := range token {
			vocabulary[t]++
		}
	}
}

func mergePairs(tokens [][]string, pair string) [][]string {
	updatedTokens := make([][]string, 0, len(tokens))
	for _, token := range tokens {
		mergedToken := make([]string, 0, len(token))
		i := 0
		for i < len(token) {
			if i < len(token)-1 && token[i]+token[i+1] == pair {
				mergedToken = append(mergedToken, pair)
				i += 2 // skip 2 position
			} else {
				mergedToken = append(mergedToken, token[i])
				i += 1 // skip 1 position
			}
		}
		updatedTokens = append(updatedTokens, mergedToken)
	}
	return updatedTokens
}

func main() {

	corpus := "low lower lowest"

	// start with character as tokens
	// each words -> for each word : append(_) -> characters, (low) -> ([l, o, w, _])

	// iterate over the corpus and create the initial vocabulary with each characters in the corpus
	words := splitTextIntoWords(corpus)
	tokens := getInitialTokens(words)
	fmt.Println(tokens)
	updateVocabulary(tokens)
	fmt.Println(vocabulary)

	maxIteration := 2
	for i := 0; i < maxIteration; i++ {
		pairCounts := countPairs(tokens)
		if len(pairCounts) == 0 {
			break
		}
		_, best_pair := findMostFreqPairs(pairCounts)
		tokens = mergePairs(tokens, best_pair)

		updateVocabulary(tokens)
	}

	fmt.Println("Final tokenization:")
	for i, token := range tokens {
		fmt.Printf("%s -> %v\n", words[i], token)
	}

}
