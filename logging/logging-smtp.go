package main

import (
	"log"
	"net/smtp"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate)
	log.Println("init started")
}

func main() {

	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()
}
