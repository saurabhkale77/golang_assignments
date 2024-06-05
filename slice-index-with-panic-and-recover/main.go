package main

import "fmt"

func main() {
	var index int
	var slice []int = []int{1, 2, 3, 4, 5}

	fmt.Printf("\nProgram to test panic and recover\n")

	fmt.Printf("\nExisting slice - %v\n", slice)

	fmt.Printf("\nPlease enter an index to retrieve the corresponding value from the slice: \n")
	fmt.Scanln(&index)

	accessSlice(slice, index)

	fmt.Printf("\nTesting panic and recover\n")
}

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\ninternal error: %v\n", r)
		}
	}()

	fmt.Printf("\nitem %d, value %d\n", index, slice[index])

	fmt.Printf("\nAt the end of accessSlice function\n")
}
