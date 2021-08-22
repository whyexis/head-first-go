package main

import "fmt"

func myFunction() {
	panic("oh no")
	fmt.Println(recover())
}

func main() {
	fmt.Println("a")
	myFunction()
	fmt.Println("b")
}