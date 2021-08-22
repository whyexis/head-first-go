package main

import "fmt"

func otherFunction() {
	fmt.Println("c")
	fmt.Println(recover())
}

func myFunction() {
	defer otherFunction()
	panic("oh no")
	fmt.Println("d")
}

func main() {
	fmt.Println("a")
	myFunction()
	fmt.Println("b")
}