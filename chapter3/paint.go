/**
 * ------------------------------------------------------------------------
 * Paint Calculator.
 * Demonstrate the use of functions to avoid repeating code.
 * Use Printf to format text output.
 * ------------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"log"
)


var metresPerLitre float64

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	 area := width * height
	 return area / metresPerLitre, nil // Return the amount of paint, with "nil" indicating there was no error
}
 
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	metresPerLitre = 10.0
	var total float64
	amount, err := paintNeeded(4.2, 3.0)
	checkError(err)
	fmt.Printf("%.2f litres needed\n", amount)
	total += amount
	amount, err = paintNeeded(5.2, 3.5)
	checkError(err)
	fmt.Printf("%.2f litres needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f litres\n", total)
}