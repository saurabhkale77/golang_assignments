package main

import (
	"assignment1/quadrilateral-calculations/shapes"
	"fmt"
)

func main() {
	var shapeNumber byte

	fmt.Printf("\nPlease enter a number to choose shape, 1 for rectangle and 2 for square:\n")
	fmt.Scanln(&shapeNumber)

	if shapeNumber != 1 && shapeNumber != 2 {
		fmt.Printf("\nInvalid shape number")
	} else {
		if shapeNumber == 1 {
			var lengthOfRect, widthOfRect float32

			fmt.Printf("\nPlease enter length of rectangle\n")
			fmt.Scanln(&lengthOfRect)

			fmt.Printf("\nPlease enter width of rectangle\n")
			fmt.Scanln(&widthOfRect)

			if lengthOfRect < 1 || widthOfRect < 1 {
				fmt.Printf("\nInvalid values")

				return
			}

			var rectangle shapes.Rectangle
			var ptrToRectangle *shapes.Rectangle = &rectangle

			ptrToRectangle.SetLength(lengthOfRect)
			ptrToRectangle.SetWidth(widthOfRect)

			shapes.PrintAreaAndPerimeter(rectangle)
		} else {
			var sideOfSquare float32

			fmt.Printf("\nPlease enter side of square\n")
			fmt.Scanln(&sideOfSquare)

			if sideOfSquare < 1 {
				fmt.Printf("\nInvalid values")

				return
			}

			var square shapes.Square
			var ptrToSquare *shapes.Square = &square

			ptrToSquare.SetSide(sideOfSquare)

			shapes.PrintAreaAndPerimeter(square)
		}
	}
}
