package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getResponseSize(url string, ch chan int) {

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	ch <- len(body)

}

func main() {
	ch := make(chan int)

	go getResponseSize("http://verge.com", ch)
	go getResponseSize("http://cnbc.com", ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
