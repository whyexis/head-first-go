package main

import "fmt"

// YOUR CODE HERE: Declare a "rectangle" struct type.
type rectangle struct {
	length float64
	width float64
}
// YOUR CODE HERE: Define a rectangleInfo function.
func rectangleInfo(r rectangle) {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func main() {
	// YOUR CODE HERE: Create a new rectangle, and set its
	// length and width fields. Pass it to rectangleInfo.
	r := rectangle{length: 4.2, width: 2.3}
	rectangleInfo(r)
}