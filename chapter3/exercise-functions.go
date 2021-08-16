/**
 * ------------------------------------------------------------------------
 * Write down what the program output will be.
 * ------------------------------------------------------------------------
 */

package main

import "fmt"
 
 func functionA(a int, b int) {
	 fmt.Println(a + b)
 }

 func functionB(a int, b int) {
	 fmt.Println(a * b)
 }

 func functionC(a bool) {
	 fmt.Println(!a)
 }

 func functionD(a string, b int) {
	 for i := 0; i < b; i ++ {
		 fmt.Print(a)
	 }
	 fmt.Println()
 }
  
 func main() {
	 functionA(2, 3)		// 5
	 functionB(2, 3)		// 6
	 functionC(true)		// false
	 functionD("$", 4)		// $$$$
	 functionA(5, 6)		// 11
	 functionB(5, 6)		// 30
	 functionC(false)		// true
	 functionD("ha", 3)		// hahaha	 
 }