package main

import "fmt"

// chan<- reciever channel & <-chan sender channel
func ping(pings chan<- string, msg string) {

	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings

	pongs <- msg
}

func main() {
	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, "PING...PONG...")
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}
