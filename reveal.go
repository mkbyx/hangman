package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
func revealRandomLetters(word string, n int) string {

	revealed := make ([]rune, len(word))
	for i := range revealed {
		revealed[i] = '_'
	}
	revealedIndices := make(map[int]bool)

	for len(revealedIndices) < n {
		index := rand.Intn(len(word))
		if !revealedIndices[index] {
			revealedIndices[index] = true
			revealed[index] = rune(word[index])
		}
	}
	return string (revealed)

}

func main() {
	var word string
	fmt.Print("Enter a word")
	fmt.Scanln(&word)
	n := len(word)/ 2 - 1
	if n < 1 {
n = 1		
	}
	revealedWord := revealRandomLetters(word,n)
	fmt.Printf("Revealed word : %\n"), revealedWord
}