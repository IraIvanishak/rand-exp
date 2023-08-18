package text

import (
	"math/rand"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generates a random text string of a specified length using
// characters from the letters slice.
func RandomText(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(letters[rand.Intn(len(letters))])
	}
	return builder.String()
}
