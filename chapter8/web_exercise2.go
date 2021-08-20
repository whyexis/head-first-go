package main

import "fmt"

type rectangle struct {
	length float64
	width float64
}

func makeSquare (r *rectangle) {
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
	r1 := rectangle{length: 4.2, width: 2.3}
	makeSquare(&r1)
	rectangleInfo(r1)

	r2 := rectangle{length: 4.5, width: 9.0}
	makeSquare(&r2)
	rectangleInfo(r2)
}