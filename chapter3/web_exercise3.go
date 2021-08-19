/*
Update your perimeter function from the previous exercise to return an error if either the length or width are negative.

    If the given length is negative, return a perimeter of 0 and an error with the message “a length of [length] is invalid”.
    If the given width is negative, return a perimeter of 0 and an error with the message “a width of [width] is invalid”.
    Otherwise, return the calculated perimeter, and an error value of nil.

In your main function, test returning an error by calling perimeter with a negative value. If the returned error value is not nil, print the error message and exit the program. Otherwise, print the returned perimeter.
 */

package main

import (
	"fmt"
	"log"
)

// perimeter takes two float64 parameters, length and width.
// And returns a float64 value representing the perimeter length.
func perimeter(length float64, width float64) (float64, error) {
	if length < 0 {
		return 0, fmt.Errorf("a length of %0.2f is invalid", length)
	}
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	return (length + width) * 2, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	total, err := perimeter(8.2, -10.0)
	checkError(err)
	plot2, err := perimeter(5.0, 5.2)
	checkError(err)
	plot3, err := perimeter(6.2, 4.5)
	checkError(err)
	total += plot2 + plot3
	fmt.Println("You'll need", total, "meters of fencing")
}