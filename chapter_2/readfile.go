package main

import (
	"fmt"
	"os"
)

func readFile() string {
	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}

	return string(data)
}
