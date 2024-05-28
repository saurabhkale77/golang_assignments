package main

import "fmt"

func main() {
	var array [8]string = [8]string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	fmt.Println("\nProgram to show slices of the array based on indexes you enter")
	fmt.Printf("\nArray is : %v\n", array)

	var firstIndex, secondIndex int

	fmt.Println("\nPlease enter the value of first index: ")
	fmt.Scanln(&firstIndex)
	fmt.Println("\nPlease enter the value of second index: ")
	fmt.Scanln(&secondIndex)

	if 0 <= firstIndex && firstIndex <= secondIndex && secondIndex <= len(array)-1 {
		fmt.Printf("\nSlice of the array from the start of the array to index %d is: ", firstIndex)
		fmt.Printf("\n%v", array[:firstIndex+1])

		fmt.Printf("\n\nSlice of the array from index %d to index %d is: ", firstIndex, secondIndex)
		fmt.Printf("\n%v", array[firstIndex:secondIndex+1])

		fmt.Printf("\n\nSlice of the array from index %d to the end of the array is: ", secondIndex)
		fmt.Printf("\n%v\n\n", array[secondIndex:])
	} else {
		fmt.Println("\nIndexes you entered are incorrect")
	}
}
