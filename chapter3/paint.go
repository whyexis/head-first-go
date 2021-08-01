/**
 * ------------------------------------------------------------------------
 * Paint Calculator.
 * Demonstrate the use of functions to avoid repeating code.
 * Use Printf to format text output.
 * ------------------------------------------------------------------------
 */

package main

import "fmt"

var metresPerLitre float64

func paintNeeded(width float64, height float64) float64 {
	 area := width * height
	 return area / metresPerLitre
}
 
func main() {
	metresPerLitre = 10.0
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f litres needed\n", amount)
	total += amount
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f litres needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f litres\n", total)
}