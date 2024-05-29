package main

import (
	"fmt"
	"reflect"
)

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
	fmt.Printf("\nThis is a value of type %s, %v", reflect.TypeOf(value).Name(), value)
}
