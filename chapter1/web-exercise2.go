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
	pebbleWeight := 0.1
	rockWeight := 1.2
	boulderWeight := 502.4
	totalWeight := pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}