package main

import "fmt"

func main() {

	// var (
	// 	name     string = "Name"
	// 	age      uint   = 27
	// 	location string = "Chicago"
	// )

	// var (
	// 	name     = "Name"
	// 	age      = 27
	// 	location = "Chicago"
	// )

	// var name = "Name"
	// var age = 27
	// var location = "Chicago"

	name, age, location := "Name", 27, "Chicago"

	fmt.Printf("%s (%d) from %s\n", name, age, location)

	action := func() {
		fmt.Println("Action as a variable")
	}

	action()

	actionWithParam := func(name string) {
		fmt.Println("Hello,", name, "!")
	}

	actionWithParam("John")
}
