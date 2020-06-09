package main

import "log"

func init() {
	log.SetPrefix("LOG : ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("init strted")
}

func main() {

	log.Println("main started")

	log.Fatalln("fatal message")

	log.Panicln("panic message")
}
