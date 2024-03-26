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
		radius := accept_input("Enter radius of circle")
		fmt.Println(circle_area(radius))
	case 2:
		len := accept_input("Enter length of rectangle")
		bth := accept_input("Enter breadth of rectangle")
		fmt.Println(rect_area(len, bth))
	case 3:
		len := accept_input("Enter length og sqaure")
		fmt.Println(square_area(len))
	}
}

func accept_input(message string) int {
	var temp int
	fmt.Println(message)
	fmt.Scan(&temp)
	return temp
}

func circle_area(radius int) float64 {
	area := 2 * math.Pi * float64(radius) * float64(radius)
	return area
}

func rect_area(len, bth int) (area int) {
	area = len * bth
	return area
}

func square_area(len int) int {
	area := len * len
	return area
}
