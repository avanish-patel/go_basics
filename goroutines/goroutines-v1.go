package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getResponseSize(url string) {

	fmt.Printf("Step 1 : %s\n", url)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Printf("Step 2 : %s\n", url)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Step 3 : Response length : %d for %s\n", len(body), url)
}

func main() {

	go getResponseSize("http://www.golang.org")
	go getResponseSize("http://www.facebook.com")
	go getResponseSize("http://www.tesla.com")

	time.Sleep(10 * time.Second)
}
