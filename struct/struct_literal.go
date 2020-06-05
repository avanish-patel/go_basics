package main

import "fmt"


type Rectangle struct {
    length int
    breadth int
    color string
}

func main() {

    var r = Rectangle{10,30,"Green"}
    fmt.Println(r)

    var r1 = Rectangle{length: 30, breadth: 10, color: "Yellow"}
    fmt.Println(r1)
}
