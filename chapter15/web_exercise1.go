package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// rollDie simulates rolling a six-sided die.
func rollDie() int {
	return rand.Intn(6) + 1
}

func rollHandler(writer http.ResponseWriter, request *http.Request) {
	// YOUR CODE HERE: Call rollDie, convert the return
	// value to a string, and convert the string to a
	// slice of bytes. Store the result in a "body"
	// variable.
	
	_, err := writer.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	// YOUR CODE HERE: Have all requests for a URL with a
	// path of "/roll" go to the rollHandler function.
	
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}