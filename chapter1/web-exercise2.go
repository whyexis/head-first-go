/**
 * ------------------------------------------------------------------------
 * Code Rework
 * Several of the variable names used in this example break Go naming conventions.
 * The code is also somewhat longer than it needs to be.
 * See if you can modify so that it’s shorter and follows conventions better, but still produces the same output.
 * Output = 503.7
 * ------------------------------------------------------------------------
 */

package main

import (
	"fmt"
)

func main() {
	var pebbleweight float64 = 0.1
	var rockweight float64 = 1.2
	var boulderweight float64 = 502.4
	var total_weight float64 = pebbleweight + rockweight + boulderweight
	fmt.Println(total_weight)
}