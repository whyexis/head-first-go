/**
 * ------------------------------------------------------------------------
 * Pool Puzzle.
 * Take code snippets and place them into the blank lines in the code.
 * Don't use the same snippet more than once, and you won't need to use all the snippets.
 * Make code that will run and produce the output shown.
 * ------------------------------------------------------------------------
 */

/**
 * ------------------------------------------------------------------------
 * Snippets.
 * err
 * error
 * errors
 * float64
 * quotient
 * divide
 * divisor
 * nil
 * ------------------------------------------------------------------------
 */

 /**
 * ------------------------------------------------------------------------
 * Output.
 * can't divide by 0
 * ------------------------------------------------------------------------
 */

 package main

 import (
	 "errors"
	 "fmt"
 )

 func divide(dividend float64, divisor float64) (float64, _______) {
	 if divisor == 0.0 {
		 return 0, _______.New("can't divide by 0")
	 }
	 return dividend / divisor, _______
 }

 func main() {
	 _______, _______ := divide(5.6, 0.0)
	 if err!= nil {
		 fmt.Println(err)
	 } else {
		 fmt.Printf("%0.2f\n", quotient)
	 }
 }