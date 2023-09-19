package main

import "fmt"

type Base struct {
	num int
}

// type "base" has a method called describe()
func (b Base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

// In cases where we want a type to have the fields from another type
// Instead of creating a variable and assigning that type to it
// We can just enter the type name and all the fields and properties will automatically be added to the parent type
type Container struct {
	Base
	str string
}

// A normal "fmt.Println(container)" will give us an output of "{{2} Hey there}"
// If we want to have a custom print thing we can create a function that outputs that
func (c Container) String() string {
	return fmt.Sprintf("The Container has a num of %v", c.num)
}

func init() {
	base := Base{num: 2}
	fmt.Println(base.describe()) // Here we are calling the method attached to base type

	container := Container{Base: base, str: "Hey there"}
	// The line below will output "The Container has a num of 2"
	fmt.Println(container)
	fmt.Println(container.num)        // We can do it this instead of container.base.num
	fmt.Println(container.describe()) // We can do it this instead of container.base.describe
}
