package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
}

func main() {
	sizes := make(chan int)
	go responseSize("https://example.com", sizes)
	go responseSize("https://golang.org", sizes)
	go responseSize("https://golang.org/doc", sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)	
}