package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("\nProgram to reverse a string using a go routine\n\n")

	var str string
	var wg sync.WaitGroup

	fmt.Printf("\nPlease enter a string\n")
	fmt.Scanln(&str)

	wg.Add(1)

	go reverseString(str, &wg)

	wg.Wait()
}

func reverseString(str string, wg *sync.WaitGroup) {
	var reverseStr string

	for _, charOfStr := range str {
		reverseStr = string(charOfStr) + reverseStr
	}

	fmt.Printf("\nReversed string is : %s\n", reverseStr)
	fmt.Printf("\nNumber of goroutines launched = %d\n", runtime.NumGoroutine())

	wg.Done()
}
