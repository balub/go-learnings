package newPackage

import "fmt"

var Array [3]string
var Slice []int

// We can't add values to a null map like "var Map = map[string]bool"
// so we need to use the make() method which will initialize a nil map into which values can be written into
var Map = make(map[string]bool)

func init() {
	Array[0] = "Hey"
	Array[1] = "There"
	Array[2] = "Traveler"
	fmt.Println(Array)
	newArray := [2]bool{true, false}
	fmt.Println(newArray)

	// With Slices you dont technically assign values to one, you just create a new slice and assign values to it
	Slice = append(Slice, 1, 2, 3, 4)
	fmt.Println(Slice)
	newSlice := []bool{true, false, false, true}
	newSlice = append(newSlice, true, true) // we create a new slice with the same name and assign previous values to it
	fmt.Println(newSlice)

	Map["isTrue"] = true
	Map["isFalse"] = false
	fmt.Println(Map)
	newMap := map[string]int{"http": 80, "https": 443}
	fmt.Println(newMap)
}
