package main

import "fmt"

func myFunction() {
	defer fmt.Println("b")
	panic("oh no")
}

func main() {
	fmt.Println("a")
	myFunction()
}