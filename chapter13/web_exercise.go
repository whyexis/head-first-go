package main

import "fmt"

func a(done chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	done <- true

}

func b(done chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	done <- true
}

func main() {
	// Create the channel
	done := make(chan bool)
	// Run each function as a goroutine,
	// passing the channel to each.
	go a(done)
	go b(done)
	// Read from the channel, which will block until the first goroutine is done.
	<-done
	// Block until the second goroutine is done
	<-done
	fmt.Println("end main()")
}
