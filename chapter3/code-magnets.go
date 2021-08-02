/**
 * ------------------------------------------------------------------------
 * Code Magnets Exercise.
 * A Go program that uses a pointer variable.
 * Reconstruct the code snippets to make a working program that will produce the given output.
 * The program should declare myInt as an integer variable,
 * and myIntPointer as a variable that holds an integer pointer.
 * Then it should assign a value to myInt, and assign a pointer to myInt as the value of myIntPointer.
 * Finally, it should print the value at myIntPointer.
 * ------------------------------------------------------------------------
 */

 package main

 import "fmt"

 func main() {
	 var myInt int
	 var myIntPointer *int
	 myInt = 42
	 myIntPointer = &myInt
	 fmt.Println(*myIntPointer)
 }