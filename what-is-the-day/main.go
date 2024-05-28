package main

import "fmt"

var dayNumberVsDayNameMap map[int]string = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
}

func main() {
	fmt.Println("\nProgram to show weekday against a number")
	fmt.Println("\nPlease enter a number : ")
	var dayNumber int
	fmt.Scanln(&dayNumber)

	dayName, isValidDayNumber := dayNumberVsDayNameMap[dayNumber]

	if !isValidDayNumber {
		fmt.Println("\nNot a day")
	} else {
		fmt.Printf("\nIt is %s", dayName)
	}
}
