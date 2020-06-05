package main

import "fmt"

type Geometry interface {
    Edges() int
}

type Pentagon struct{}
type Hexagon struct{}
type Octagon struct{}


func (p Pentagon) Edges() int {
    return 5
}

func (h Hexagon) Edges() int {
    return 6
}

func (o Octagon) Edges() int {
    return 8
}


func Parameter(g Geometry, v int) int {
    return g.Edges() * v
}


func main() {
    p := new(Pentagon)
    h := new(Hexagon)
    o := new(Octagon)

    g := []Geometry{p,h,o}

    for _, i := range g {
        fmt.Println(Parameter(i, 5))
    }
}
