package text

import (
	"math/rand"
	"strings"
)

func RandomSentence(minWords, maxWords int) string {
	wordCount := rand.Intn(maxWords-minWords+1) + minWords
	words := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		words[i] = RandomText(rand.Intn(10) + 1)
	}
	return strings.Join(words, " ") + "."
}
