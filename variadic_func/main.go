package main

import (
	"fmt"
	"reflect"
)

func main() {

	printColors("red", "green", "blue", "yellow")

	fmt.Println(calculateArea("RECTANGLE", 20, 30))
	fmt.Println(calculateArea("SQUARE", 20))

	variadicArgs(1, 2.0, "Red",
		[]string{"red", "green", "orange"},
		map[string]int{"apple": 1, "banana": 2})
}

func printColors(s ...string) {

	for _, v := range s {
		fmt.Println(v)
	}
}

func calculateArea(t string, d ...int) int {

	area := 1

	for _, value := range d {
		if t == "RECTANGLE" {
			area *= value
		} else if t == "SQUARE" {
			area = value * value
		}
	}
	return area

}

func variadicArgs(i ...interface{}) {

	for _, v := range i {
		fmt.Println(v, "--", reflect.TypeOf(v).Kind())
	}
}
