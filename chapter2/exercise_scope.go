/**
 * ------------------------------------------------------------------------
 * Some of the lines of code below will result in a compile error,
 * because they refer to a variable that is out of scope.
 * Cross out the lines that have errors.
 * ------------------------------------------------------------------------
 */

package main

import (
	"fmt"
)

 func main() {
	 //a = "a"
	 b := "b"
	 if true {
		 c := "c"
		 if true {
			 d := "d"
			 // fmt.Println(a)
			 fmt.Println(b)
			 fmt.Println(c)
			 fmt.Println(d)
		 }
		 // fmt.Println(a)
		 fmt.Println(b)
		 fmt.Println(c)
		 // fmt.Println(d)
	 }
	 // fmt.Println(a)
	 fmt.Println(b)
	 // fmt.Println(c)
	 // fmt.Println(d)
 }