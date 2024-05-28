package main

import (
	stringtoslice "assignment1/word-count/string-to-slice"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nProgram to show words with highest frequency in a string")
	fmt.Println("\nPlease enter a string containing words separated by space: ")

	var inputReader = bufio.NewReader(os.Stdin)
	// specify till 1st occurence of which chatacter we have to read, here it is \n
	// read string
	strContainingWords, err := inputReader.ReadString('\n')

	// if any error occurs while reading the string
	if err != nil {
		fmt.Println("\nError in reading the input")
		return
	}

	// remove appended \n at the end
	strContainingWords = strings.TrimSpace(strContainingWords)

	// if entered string is empty
	if strContainingWords == "" {
		fmt.Println("\nEntered string is empty")
		return
	}

	// GetSliceContainingWordsWithHighestFrequencyFor the entered string
	finalSliceContainingWordsWithHighFrequency, maxFrequencyOfAnyWord := stringtoslice.GetSliceContainingWordsWithHighestFrequencyForStr(strContainingWords)
	fmt.Printf("\nList of words with high frequency(%d) is - \n%v", maxFrequencyOfAnyWord, finalSliceContainingWordsWithHighFrequency)
}
