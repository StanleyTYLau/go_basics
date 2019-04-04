package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

// Define structs
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	length, width float64
}

// Define methods
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Hello World!")
	// circle := Circle{0, 0, 5}
	// rectangle := Rectangle{10, 5}

	var circle Shape = Circle{0, 0, 5}
	var rectangle Shape = Rectangle{10, 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

	fmt.Printf("Rectangle Area: %f\n", rectangle.area())
	fmt.Printf("%T\n", circle)

}
