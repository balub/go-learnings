package main

import "fmt"

func main() {
	var myFirstName string = "Balu"
	fmt.Println((myFirstName))

	var myMiddleName = "Babu"
	fmt.Println((myMiddleName))

	myLastName := "Naidu"
	fmt.Println((myLastName))

	firstName, middleName, lastName := "Balu", "Babu", "Naidu"
	fmt.Println(firstName, middleName, lastName)

	first, _, last := "Balu", "Babu", "Naidu" // we use _ to ignore the particular value
	fmt.Println(first, last)

	var (
		saishoNAME = "Balu"
		saigoNAME  = "Babu"
	)
	fmt.Println(saishoNAME, saigoNAME)

}
