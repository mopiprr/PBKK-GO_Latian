package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to reverse a single word
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function to reverse each word in a string
func reverseWords(input string) string {
	words := strings.Fields(input)
	if len(words) < 3 {
		return "Please enter at least 3 words."
	}

	for i, word := range words {
		words[i] = reverseWord(word)
	}
	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string with at least 3 words: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := reverseWords(input)
	fmt.Println("Output:", result)
}
