package main

import "fmt"

type Polygons interface {
    Perimeter()
}

type Object interface {
    NumberOfSides()
}

type Pantagon int

func (p Pantagon) Perimeter() {
    fmt.Println("Perimeter of Pantagon : ", 5*p)
}

func (p Pantagon) NumberOfSides() {
    fmt.Println("Pantagon has 5 side")
}


func main() {

    var p Polygons = Pantagon(50)
    p.Perimeter()

    var o Pantagon = p.(Pantagon)
    o.NumberOfSides()

    var j Object = Pantagon(10)
    j.NumberOfSides()

    var k Pantagon = j.(Pantagon)
    k.Perimeter()

}
