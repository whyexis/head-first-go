package main

import "fmt"

// greeting takes a channel as a parameter
func greeting(myChannel chan string) {
	myChannel <- "hi" // Send a value over the channel
}

func main() {
	myChannel := make(chan string) // Create a new channel
	go greeting(myChannel)         // Pass the channel to function running in a new goroutine
	//fmt.Println(<-myChannel)       // Receive a value from the channel
	receivedValue := <-myChannel // Could also store the received value in a variable instead
	fmt.Println(receivedValue)
}
