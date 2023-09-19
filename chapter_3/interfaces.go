package main

import (
	"fmt"
	"math"
)

type Dimension float32

type Rectangle struct {
	Length, Breadth Dimension
}

type Circle struct {
	Radius Dimension
}

// Interfaces in Go are similiar to interfaces in Java
// They define the type of the functions that need to be implemented in any type that implements them
type Shapes interface {
	Area() Dimension
	Perimeter() Dimension
}

func (r Rectangle) Area() Dimension {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() Dimension {
	return 2 * (r.Length + r.Breadth)
}

func (c Circle) Area() Dimension {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() Dimension {
	return 2 * math.Pi * c.Radius
}

// In this case this function accepts a generic type that implements the Shape interface
// Since both the "Rectangle" and "Circle" types implement the functions defined in the interface
// They can be passed into this 
func Measure(s Shapes) {
	fmt.Println(s)
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func init() {
	rectangle := Rectangle{Length: 3, Breadth: 3}
	circle := Circle{Radius: 4}
	// In this case we crea
	Measure(rectangle)
	Measure(circle)

}
