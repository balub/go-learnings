package main

import (
	"fmt"
	"os"
)

func writeToFile(content string) {
	err := os.WriteFile("writeFile.txt", []byte(content+"Added in writeToFile function"), 0666)
	if err != nil {
		fmt.Println(err)
	}
}
