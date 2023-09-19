package main

import "fmt"

func main() {
	fmt.Println("Hello from main.go")
	fileContent := readFile()
	fmt.Println(fileContent)
	writeToFile(fileContent)
}
