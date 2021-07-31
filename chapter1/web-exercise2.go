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

// Since we're only importing one package, we can use the
// single-line format for "import".
import "fmt"

func main() {
    // We can use short variable declarations for everything.
	// Later words in multi-word variable names should be capitalized.
	pebbleWeight := 0.1
	rockWeight := 1.2
	boulderWeight := 502.4
    // Underscores are legal in variable names, but are not conventional.
	// Remove underscore and capitalize the second word.
	totalWeight := pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}