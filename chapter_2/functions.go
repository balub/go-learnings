package main

import "fmt"

// Normal function, with input and output types defined
// In the signature the name followed by the type
func add(a int, b int) int {
	return a + b
}

// If all input types are same we can just define them once in the end as seen below
func subtract(a, b int) int {
	return a - b
}

// In Go functions are return 2 values, Go does not have a concept of objects
// The 2 values being returned are full blown variables
// So in the function below we are not returning an object with 2 values of type int
func addAndSubtract(a, b int) (int, int) {
	return add(a, b), subtract(a, b)
}

// The function below is the same as the one above, only difference is
// We are naming the output variables and these names are essentially locally scoped variables to the function
// So we assign a value to them and just have to say return nothing more
// Because these local variables are initialized in the return type we dont need to use the ":=" syntax
// when assigning a value to them inside the function
func addAndSubtractWithName(a, b int) (AdditionResult int, SubtractionResult int) {
	AdditionResult = add(a, b)
	SubtractionResult = subtract(a, b)
	return
}

func init() {
	res_sum, res_diff := addAndSubtract(1, 2)
	// PrintF is like how you do template strings in Go, here %v is a generic catch all for any value type
	// but if we wanted to print integers only we'd use %i checkout "verbs" in the std library
	fmt.Printf("The Sum is %v and the difference is %v\n", res_sum, res_diff)

	// In the case below, I only want the result of the addition and I dont care about the other values being returned
	// So I use "_" to basically ignore that value, Go's garbage collection then removes that variables during runtime
	result, _ := addAndSubtractWithName(2, 3)
	fmt.Println(result)

}
