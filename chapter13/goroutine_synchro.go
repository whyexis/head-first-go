package main

import (
	"fmt"
	"time"
)

// reportNap takes the name of the sleeping goroutine
// and the time to sleep for in seconds
func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a" // will block on this send while "main" is still asleep
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}