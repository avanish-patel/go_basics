package main

import "fmt"


type Book struct {
    author, title string
}

type Magazine struct {
    title string
    issue int
}

func (b *Book) Assign(n,t string) {
    b.author = n
    b.title = t
}

func (b *Book) Print() {
    fmt.Printf("Author: %s, Title: %s\n", b.author, b.title)
}

func (m *Magazine) Assign(n string, t int) {
    m.title = n
    m.issue = t
}

func (m *Magazine) Print() {
    fmt.Printf("Title: %s, Issue: %d\n", m.title,m.issue)
}


type Printer interface {
    Print()
}


func main() {

    var b Book
    var m Magazine

    b.Assign("Jack Rabbit","Book of Rabbits")
    m.Assign("Rabbit Weekly",26)

    var i Printer

    i = &b
    i.Print()

    i = &m
    i.Print()
}
