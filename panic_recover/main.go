package main

import "fmt"

func main() {

	var action int

	// fmt.Println("Enter 1 : student and 2: professional")
	// fmt.Scanln(&action)

	// switch action {
	// case 1:
	// 	fmt.Println("You're Student!")
	// case 2:
	// 	fmt.Println("You're Professional!")
	// default:
	// 	panic(fmt.Sprintf("You're %d", action))
	// }
	// fmt.Println()
	fmt.Println("Enter 1: US and 2:UK")
	fmt.Scanln(&action)
	switch action {
	case 1:
		fmt.Println("US")
	case 2:
		fmt.Println("UK")
	default:
		panic(fmt.Sprintf("You're from Marsh : %d", action))
	}

	defer func() {
		action := recover()
		fmt.Println(action)
	}()
}
