package stringtoslice

import "strings"

func GetSliceContainingWordsWithHighestFrequencyForStr(str string) ([]string, int) {
	// split str by space
	sliceOfSplittedStr := strings.Split(str, " ")

	// to maintain mapping of words with their number of occurences
	var wordVsCountMap map[string]int = map[string]int{}

	// count occurences for each word
	for _, word := range sliceOfSplittedStr {
		wordVsCountMap[word]++
	}

	// find out maximum occurence of any word
	maxFrequencyOfAnyWord := 0
	for _, frequencyOfWord := range wordVsCountMap {
		if frequencyOfWord > maxFrequencyOfAnyWord {
			maxFrequencyOfAnyWord = frequencyOfWord
		}
	}

	// arrange maximum occured words sequentially without repetition
	var finalSliceContainingWordsWithHighFrequency []string
	for _, word := range sliceOfSplittedStr {
		if wordVsCountMap[word] == maxFrequencyOfAnyWord && !isWordAlreadyExistsInSlice(word, finalSliceContainingWordsWithHighFrequency) {
			finalSliceContainingWordsWithHighFrequency = append(finalSliceContainingWordsWithHighFrequency, word)
		}
	}

	// return the final slice
	return finalSliceContainingWordsWithHighFrequency, maxFrequencyOfAnyWord
}

func isWordAlreadyExistsInSlice(wordToBeMatched string, finalSliceContainingWordsWithHighFrequency []string) bool {
	// this function is to avoid repetition of elements in final array
	for _, word := range finalSliceContainingWordsWithHighFrequency {
		if word == wordToBeMatched {
			return true
		}
	}

	return false
}
