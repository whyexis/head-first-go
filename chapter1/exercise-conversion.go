/**
 * ------------------------------------------------------------------------
 * Conversions lets you convert a value from one type to another type.
 * Provide the type you want to convert a value to, immediately followed by
 * the value you want to convert in parentheses.
 * ------------------------------------------------------------------------
 */
 
package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = price * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = price + taxRate
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= availableFunds)
}