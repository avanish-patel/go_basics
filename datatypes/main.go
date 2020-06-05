package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	goIsFun       = true
	integer       = 1<<32 - 1
	complexNumber = -1 + 2i
	name          = "Name"
)

func main() {
	const f = "%T(%v)\n"

	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, integer, integer)
	fmt.Printf(f, complexNumber, complexNumber)
	fmt.Printf(f, name, name)

	// type convesrion
	// The expression T(v) converts the value v to the type T.
	var i int = 43
	var ft float64 = float64(i)
	var u uint = uint(ft)

	fmt.Printf(f, i, i)
	fmt.Printf(f, ft, ft)
	fmt.Printf(f, u, u)

	// Integer conversion
	fmt.Println(strconv.Atoi("100"))
	fmt.Println(strconv.ParseInt("-52541", 10, 32))
	fmt.Println(strconv.ParseInt("101010101010101010", 10, 64))

	// Float conversion
	fmt.Println(strconv.ParseFloat("3.1415926535", 8))

	// Boolean convrersion from string
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("F"))
	fmt.Println(strconv.ParseBool("0"))

	// String to boolean
	bl := strconv.FormatBool(true)
	fmt.Println(bl, reflect.TypeOf(bl))

	ftc := strconv.FormatFloat(3.2334, 'E', -1, 32)
	fmt.Println(ftc, reflect.TypeOf(ft))

}
