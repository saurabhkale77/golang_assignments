package shapes

import "fmt"

type Quadrilateral interface {
	calculateArea() float32
	calculatePerimeter() float32
}

type Rectangle struct {
	length float32
	width  float32
}

type Square struct {
	side float32
}

func (rectangle Rectangle) calculateArea() float32 {
	return rectangle.length * rectangle.width
}

func (rectangle Rectangle) calculatePerimeter() float32 {
	return 2 * (rectangle.length + rectangle.width)
}

func (square Square) calculateArea() float32 {
	return square.side * square.side
}

func (square Square) calculatePerimeter() float32 {
	return 4 * square.side
}

func (rectangle *Rectangle) SetLength(length float32) {
	(*rectangle).length = length
}

func (rectangle *Rectangle) SetWidth(width float32) {
	(*rectangle).width = width
}

func (square *Square) SetSide(side float32) {
	(*square).side = side
}

func PrintAreaAndPerimeter(quadrilateral Quadrilateral) {
	fmt.Printf("\nArea is : %.2f", quadrilateral.calculateArea())
	fmt.Printf("\nPerimeter is : %.2f", quadrilateral.calculatePerimeter())
}
