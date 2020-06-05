package main

import "fmt"

const (
	StatusOk       = 200
	StatusCreated  = 201
	StatusAccepted = 202
)

func main() {

	const Pi = 3.14
	const True = true

	fmt.Println(StatusOk, StatusCreated, StatusAccepted)
	fmt.Println(Pi, True)
}
