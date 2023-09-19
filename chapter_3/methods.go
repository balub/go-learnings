package main

// import "fmt"

// // Consider this TS type definition
// // type Dimension = number;

// // interface Rectangle{
// // 	Length: Dimension,
// // 	Breadth: Dimension
// // 	Area: func() Dimension
// // 	Perimeter: func() Dimension
// // }

// // We create custom types in Go in the below format
// // We capitalize the first letter to make it exportable
// type Dimension int

// // We create custom types similar to interfaces in JS in Go in the below format
// type Rectangle struct {
// 	Length, Breadth Dimension
// }

// // We can add custom methods to created methods by creating functions in the below format
// // These methods are "attached" to the type we create
// func (r Rectangle) Area() Dimension {
// 	return r.Length * r.Breadth
// }

// func (r Rectangle) Perimeter() Dimension {
// 	return 2 * (r.Length + r.Breadth)
// }

// func init() {
// 	// We can assign values directly in the constructor
// 	// Constructor is the {}
// 	rectangle := Rectangle{Length: 3, Breadth: 3}
// 	fmt.Println(rectangle)
// 	fmt.Println(rectangle.Area()) // Area method attached to variable rectangle of type Rectangle
// 	fmt.Println(rectangle.Perimeter())
// 	// If no values are assigned by default the values are nil, or zero if int type
// 	rect2 := Rectangle{}
// 	fmt.Println(rect2)
// 	// Values can also be assigned like this
// 	rect3 := Rectangle{}
// 	rect3.Length = 4
// 	rect3.Breadth = 3
// 	fmt.Println(rect3)

// }
