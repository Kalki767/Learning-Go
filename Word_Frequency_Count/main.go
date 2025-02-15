package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// a function to accept string input from the user by reading the whole line and moving to the next line
func stringInput() string {
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')

	return strings.TrimSpace(inp)
}

// a function to count the frequency of each word in a sentence by ignoring the punctuation and case sensitivity
func FrequencyCount(sentence string) map[string]int{
	frequency := make(map[string]int)
	splitted_sentence := strings.Fields(sentence)
	punctuation := " .,!?;:\"'()[]{}<>"
	

	for _,word := range splitted_sentence{
		cleaned_word := strings.Trim(word,punctuation)
		lower := strings.ToLower(cleaned_word)
		frequency[lower] += 1
	}
	return frequency
}

// main function to accept the string from the user and call the FrequencyCount function
func main(){
	fmt.Print("Please Enter the string to be counted: ")
	sentence := stringInput()
	frequency := FrequencyCount(sentence)
	fmt.Println(frequency)
}