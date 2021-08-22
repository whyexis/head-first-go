package main

import "fmt"

func main() {
	fmt.Println("a")
	panic("oh no")
	fmt.Println("b")
}

/* Ouptut Prediction
 * a
 * oh no
 */