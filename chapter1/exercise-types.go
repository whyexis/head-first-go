/**
 * ------------------------------------------------------------------------
 * Go is statically typed.
 * If you use the wrong type of value in the wrong place, Go will let you know.
 * ------------------------------------------------------------------------
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("42\t\t", reflect.TypeOf(42)) // int
	fmt.Println("3.1415\t\t", reflect.TypeOf(3.1415)) // float64
	fmt.Println("true\t\t", reflect.TypeOf(true)) // bool
	fmt.Println("Hello, Go!\t", reflect.TypeOf("Hello, Go!")) // string
	fmt.Println("25\t\t", reflect.TypeOf(25)) // int
	fmt.Println("true\t\t", reflect.TypeOf(true)) // bool
	fmt.Println("5.2\t\t", reflect.TypeOf(5.2)) // float64
	fmt.Println("1\t\t", reflect.TypeOf(1)) // int
	fmt.Println("false\t\t", reflect.TypeOf(false)) // bool
	fmt.Println("1.0\t\t", reflect.TypeOf(1.0)) // float64
	fmt.Println("hello\t\t", reflect.TypeOf("hello")) // string
}