package main

import (
	"fmt"
	"errors"
)

// Divide function attempts to divide two integers.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// In Go, we must check for errors explicitly in your code.
// What that means is that if we have a function that can error out
// We must explicitly pass a value for error, and check for it in the place it is being called
// Usually if error value exists we return the error value, if an error does not exist we return a nil
// So when the function is called if err == nil it means no errors were returned
// if err != nil , it means errors were returned
func init() {
	// Example 1: Divide two numbers successfully.
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 2: Attempt to divide by zero.
	result, err = Divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
