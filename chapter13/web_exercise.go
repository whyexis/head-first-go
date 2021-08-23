package main

import "fmt"

func a(channel chan string) {
	for i := 0; i < 50; i++ {
		channel <- "a"
	}
}

func b(channel chan string) {
	for i := 0; i < 50; i++ {
		channel <- "b"
	}
}

func main() {
	channel := make(chan string)
	go a(channel)
	go b(channel)
	for i := 0; i < 50; i++ {
		fmt.Printf(<-channel)
		fmt.Printf(<-channel)
	}
	fmt.Println("end main()")
}
