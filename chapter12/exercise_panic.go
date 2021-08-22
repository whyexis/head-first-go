package main

import "fmt"

func snack() {
	defer fmt.Println("Closing refrigerator")
	fmt.Println("Opening refrigerator")
	panic("refrigerator is empty")
}

func main() {
	snack()
}

/* Output
 * Opening refrigerator
 * Closing refrigerator
 * panic: refrigerator is empty
 * go routine 1 [running]:
 * main.snack()
 *	<path>/main.go: 8 <address>
 * main.main()
 	<path>/main.go: 12 <address>
 */