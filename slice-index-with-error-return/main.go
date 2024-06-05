package main

import (
	"errors"
	"fmt"
)

func main() {
	var index int
	var slice []int = []int{1, 2, 3, 4, 5}

	fmt.Printf("\nProgram to return error if invalid index value\n")

	fmt.Printf("\nExisting slice - %v\n", slice)

	fmt.Printf("\nPlease enter an index to retrieve the corresponding value from the slice: \n")
	fmt.Scanln(&index)

	valueAtIndex, err := accessSlice(slice, index)

	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\nitem %d, value %d\n", index, valueAtIndex)
	}
}

var indexOutOfBoundsError error = errors.New("index should be less than the length of the slice and greater than 0")

func accessSlice(slice []int, index int) (int, error) {
	if index > len(slice)-1 || index < 0 {
		return 0, indexOutOfBoundsError
	}

	return slice[index], nil
}
