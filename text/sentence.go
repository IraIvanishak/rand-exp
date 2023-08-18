package text

import (
	"math/rand"
	"strings"
)

// Generates a random sentence containing a variable number of words,
// where the word count falls within the specified range
func RandomSentence(minWords, maxWords int) string {
	wordCount := rand.Intn(maxWords-minWords+1) + minWords
	words := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		words[i] = RandomText(rand.Intn(10) + 1)
	}
	return strings.Join(words, " ") + "."
}
