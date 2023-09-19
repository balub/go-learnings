package main

import "fmt"

// We dont have to explicitly state return all the times
// But if we do this we must not state the return type for the function
func birthday(age int) {
	age++
}

// We expect this function to modify the value of the entity being passed in , in this case a variable
// Since we pass by value hence a duplicate ends up getting passed normally like in the function above
// Here we pass a pointer to the variable, denoted by prefixing the input type and the variable name with "*"
func birthdayWithPointers(age *int) {
	// *age essentially means data this pointer is pointing too
	*age++
}

func init() {
	age := 26
	fmt.Printf("My unmodified age is:%v\n", age)

	// The code below wont work because in Go in functions we can only pass by values
	// This means when we pass a value we don't pass a reference to it like say in JS
	// Instead we pass a duplicate of the original value, in the "birthday" function we want to modify the value of "age"
	// but since we pass a duplicate we original age does not change
	birthday(age)
	birthday(age)
	birthday(age)
	fmt.Printf("My New modified age is:%v\n", age) // Value of "age" will continue to be 26 not 29

	// So when we are working with functions that change the value of the data we use pointers
	// If we want to modify the value of the thing being passed in we use a pointer
	// Since we want to use pointers to modify the value we need to pass in the address of the pointer to the function
	birthdayWithPointers(&age)
	birthdayWithPointers(&age)
	birthdayWithPointers(&age)
	fmt.Printf("My Actual modified age is:%v\n", age) // Value of "age" will continue to be 26 not 29

}
