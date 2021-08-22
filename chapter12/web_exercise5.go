package main

import "fmt"

func otherFunction() {
	// This runs when the deferred function call is made.
	fmt.Println("c")
	// This recovers from the panic and prints the panic message.
	fmt.Println(recover())
}

func myFunction() {
	// This call gets made when myFunction panics.
	defer otherFunction()
	panic("oh no")
	// Execution does not resume here, because
	// this function exited due to the panic.
	fmt.Println("d")
}

func main() {
	fmt.Println("a")
	myFunction()
	// Execution resumes after recovering.
	fmt.Println("b")
}

/* Ouptut Prediction
 * a
 * c
 * oh no
 * b
 * Because the program recovered from the panic,
 * it doesn't exit early, and so no stack trace is printed.
 */