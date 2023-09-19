package main

import "fmt"

// In Go, errors are always handled in the following way
// We pass errors as a explicit value in Go

func CalculatorWithErrorDesignPattern() {

	var operation string
	var num1, num2 int

	fmt.Println("Enter operation 1.Addition 2.Subtraction 3.Multiplication 4.Division")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter First number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter Second number")
	fmt.Scanf("%d", &num2)
	switch operation {
	case "add":
		fmt.Println(add(num1, num2))
	case "sub":
		fmt.Println(subtract(num1, num2))
	case "mul":
		fmt.Println(multiply(num1, num2))
	case "div":
		// Here "divide" function returns 2 values the result and the error value
		// The result is nil if the operation errors out and error is nil if the operation is good
		res, err := Divide(num1, num2)
		// if err == nil, no errors exist
		// if err !=nil, errors exist
		if err != nil {
			fmt.Println("cannot divide numbers")
		} else {
			fmt.Println(res)
		}
	}
}

func multiply(a, b int) int {
	return a * b
}
