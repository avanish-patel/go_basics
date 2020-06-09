package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getResponseSize(url string) {

	defer wg.Done()

	fmt.Println("Step 1 " + url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Step 2 " + url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Step 3 %s, Size : %d\n", url, len(body))
}

func main() {

	wg.Add(3)

	go getResponseSize("http://google.com")
	go getResponseSize("http://facebook.com")
	go getResponseSize("https://duckduckgo.com")

	wg.Wait()
}
