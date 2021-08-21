package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

// Convert this function to a method on the "rectangle" type.
func makeSquare(r *rectangle) {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}

func rectangleInfo(r rectangle) {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func main() {
	var r rectangle
	r.length = 4.2
	r.width = 2.3
	rectangleInfo(r)
}
