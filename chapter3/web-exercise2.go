/*
You’re fed up with the neighborhood wildlife destroying your gardens. You have three rectangular garden plots that you need to surround with fencing:

    An 8.2 by 10 meter plot
    A 5 by 5.2 meter plot
    A 6.2 by 4.5 meter plot

You need to figure out how much fencing to buy. You’ll need to calculate the perimeter (two times the length plus two times the width) for each of these plots, then add the three perimeters together to get the total length of fencing.

Write a perimeter function that takes two float64 parameters representing the length and width. perimeter should return a float64 value representing the perimeter length.
*/

package main

import "fmt"

// perimeter takes two float64 parameters, length and width.
// And returns a float64 value representing the perimeter length.
func perimeter(length float64, width float64) float64 {
	return (length + width) * 2
}

func main() {
	total := perimeter(8.2, 10.0)
	total += perimeter(5.0, 5.2)
	total += perimeter(6.2, 4.5)
	fmt.Println("You'll need", total, "meters of fencing")
}