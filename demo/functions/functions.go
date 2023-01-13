package main

import "fmt"

func double(num int) int {
	return num * num
}

func add(a, b int) int { // if both the params are of same type we can just define it once
	return a + b
}

func main() {
	data := double(9)
	fmt.Println(data)
}
