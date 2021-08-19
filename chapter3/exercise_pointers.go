/**
 * ------------------------------------------------------------------------
 * Update negate function to pass a pointer instead of passing the value.
 * ------------------------------------------------------------------------
 */

package main

import "fmt"
 
 func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
 }

 func main() {
	 truth := true
	 negate(&truth)
	 fmt.Println(truth)
	 lies := false
	 negate(&lies)
	 fmt.Println(lies)
 }
 