package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

// Convert this function to a method on the "rectangle" type.
func (r *rectangle) makeSquare() {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}

func (r *rectangle) info() {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func main() {
	longRectangle := rectangle{length: 4.2, width: 2.3}
	fmt.Println("------------------------------------")
	fmt.Println("Long rectangle before makeSquare()")
	longRectangle.info()
	longRectangle.makeSquare()
	fmt.Println("------------------------------------")
	fmt.Println("Long rectangle after makeSquare()")
	longRectangle.info()

	fmt.Println("------------------------------------")

	wideRectangle := rectangle{length: 4.5, width: 9.0}
	fmt.Println("Wide rectangle before makeSquare()")
	wideRectangle.info()
	wideRectangle.makeSquare()
	fmt.Println("------------------------------------")
	fmt.Println("Wide rectangle after makeSquare()")
	wideRectangle.info()
}
