package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to take string input from the user by reading the whole line and ignoring spaces
func stringInput()string{

	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')

	return strings.TrimSpace(inp)
}

// Function that accepts string and returns a reversed one using byte slice
func reverseBytes(s string) string {
	bytes := []byte(s) // Convert string to byte slice
	n := len(bytes)

	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i] // Swap
	}

	return string(bytes) // Convert back to string
}

// Function to check if the word is a palindrome or not
func PalindromeCheck(word string) bool {
	punctuation := " .,!?;:\"'()[]{}<>"
	cleaned_word := strings.Trim(word, punctuation)
	return cleaned_word == reverseBytes(cleaned_word)
}

// Main function to take input from the user and check if the word is a palindrome or not
func main() {

	fmt.Print("Please enter a string to be checked: ")
	word := stringInput()
	ans := PalindromeCheck(word)
	if ans{
		fmt.Printf("The word %s is a palindrome", word)
	} else {
		fmt.Printf("The word %s is not a palindrome", word)
	}
}