package main

import "fmt"
// Employee is an interface for printing employee details
type Employee interface {
    PrintName(name string)
    PrintSalary(basic int, tax int) int
}


// Emp user define type
type Emp int


func (e Emp) PrintName(name string) {
    fmt.Println("Employee Id:\t", e)
    fmt.Println("Employee Name:\t",name)
}

func (e Emp) PrintSalary(basic, tax int) int {

    salary := (basic * tax)/100
    return basic - salary
}

func main() {
    var e Employee
    e = Emp(1)
    e.PrintName("John Doe")
    fmt.Println(e.PrintSalary(25000,5))
}

