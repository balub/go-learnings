package main

import "fmt"
import "chapter_1/go/io/newPackage"

// init() functions are functions that are executed before the main() function
func init() {
	fmt.Println("Init One")
}

// We can have many init functions and they all will be executed before the main() function
// However init() functions themselves are executed in the order they are defined
func init() {
	fmt.Println("Init Two")
}

// The file can be called anything for as long as its in the main package and has the main() function
func main() {
	// variable declarations
	var message string = "This is a string"
	var number int
	number = 32
	isTrue := true // initialize and as sign value to a variable

	fmt.Println("Hello! World")
	fmt.Println(message)
	fmt.Println(number)
	fmt.Println(isTrue)
	// I can call this function here because its from the same main package
	// The Go compiler kinda takes all the code from multiple files but same package and puts them into one file during execution
	// Notice how the first letter is not capitalized, this means is a private function available only in the main package
	printValue()
	// The function below is being imported from another package
	// Notice how the first letter is capitalized, this means is a public function
	newPackage.PrintPackageValue()
	// Capitalizing the first letter to export it applies to variables and constants too
	fmt.Println(newPackage.ConstFromOtherPackage)
	fmt.Printf(newPackage.VariableFromOtherPackage)
}
