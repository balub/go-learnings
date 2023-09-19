package main

import "fmt"

// In Go, errors are always handled in the following way
// We pass errors as a explicit value in Go

func Calculator() {

	var operation string
	var num1, num2 int

	fmt.Println("Enter operation 1.Addition 2.Subtraction 3.Multiplication 4.Division")
	// "%i" is the verb for integer
	// We use "&operation" because we want to pass a pointer so that we can modify the value of operation
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter First number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter Second number")
	fmt.Scanf("%d", &num2)
	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "sub":
		fmt.Println(num1 - num2)
	case "mul":
		fmt.Println(num1 * num2)
	case "div":
		fmt.Println(num1 / num2)
	}
}
