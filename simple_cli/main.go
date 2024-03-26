package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter 1 to calculate area of a circle, 2 for rectangle, 3 for square")
	var choice int

	fmt.Scan(&choice)
	switch choice {
	case 1:
		var radius int
		fmt.Println("Enter radius of circle")
		fmt.Scan(&radius)
		fmt.Println(2 * math.Pi * float64(radius) * float64(radius))
	case 2:
		var len, bth int
		fmt.Println("Enter length of rectangle")
		fmt.Scan(&len)
		fmt.Println("Enter breadth of rectangle")
		fmt.Scan(&bth)
		fmt.Println(len * bth)
	case 3:
		var len int
		fmt.Println("Enter length of square")
		fmt.Scan(&len)
		fmt.Println(len * len)
	}
}
