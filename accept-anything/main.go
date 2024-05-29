package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var numberForDataType int

	fmt.Printf("\nPlease enter a number between 1 to 4:\n")
	fmt.Scanln(&numberForDataType)

	switch numberForDataType {
	case 1:
		AcceptAnything(5)
	case 2:
		AcceptAnything("Hello")
	case 3:
		AcceptAnything(true)
	case 4:
		AcceptAnything(Person{"Jack", 24})
	default:
		fmt.Printf("\nInvalid number\n")
	}
}

func AcceptAnything(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Printf("\nThis is a value of type Integer, %d", value)
	case string:
		fmt.Printf("\nThis is a value of type String, %s", value)
	case bool:
		fmt.Printf("\nThis is a value of type Boolean, %v", value)
	case Person:
		fmt.Printf("\nThis is a value of type Person, %v", value)
	}
}
