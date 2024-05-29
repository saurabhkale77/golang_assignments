package main

import "fmt"

type Rectangle struct {
	length float32
	width  float32
}

func (rectangle Rectangle) calculateArea() float32 {
	return rectangle.length * rectangle.width
}

func (rectangle Rectangle) calculatePerimeter() float32 {
	return 2 * (rectangle.length + rectangle.width)
}

func main() {
	var rectangle Rectangle

	fmt.Printf("\nPlease enter length of rectangle\n")
	fmt.Scanln(&rectangle.length)

	fmt.Printf("\nPlease enter width of rectangle\n")
	fmt.Scanln(&rectangle.width)

	if rectangle.length < 1 || rectangle.width < 1 {
		fmt.Printf("\nInvalid values")
		return
	} else {
		fmt.Printf("\nArea of rectangle is : %.2f", rectangle.calculateArea())
		fmt.Printf("\nPerimeter of rectangle is : %.2f", rectangle.calculatePerimeter())
	}
}
