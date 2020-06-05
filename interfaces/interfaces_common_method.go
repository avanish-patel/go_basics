package main

import "fmt"

type Vehicle interface {
    Structure() []string
    Speed() string
}

type Human interface {
    Structure() []string
    Performance() string
}


type Car string


func (c Car) Structure() []string {
    return []string {"ECU","Engine","Air Filter","Wiper"}
}

func (c Car) Speed() string {
    return "100Mph"
}

type Man string 

func (m Man) Structure() []string {
    return []string{"Brain","Heart","Nose","Eyelashes","Stomach"}
}

func (m Man) Performance() string {
    return "5Mph"
}

func main() {
    var bmw Vehicle = Car("BMW")
    var man Human = Man("John Doe")

    for i, car := range bmw.Structure() {
        fmt.Printf("%-15s <=====> %15s\n", car, man.Structure()[i])
    }
}
