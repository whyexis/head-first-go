/*
 * Define a printLines function that takes a slice of strings as a parameter,
 * and prints each element of that slice on a separate line.
 * Then, in main, get a slice of the daysOfWeek array containing just the weekdays:
 * “Monday”, “Tuesday”, “Wednesday”, “Thursday”, and “Friday”.
 * Pass that slice to printLines.
 */
package main

import "fmt"

// YOUR CODE HERE: Define a printLines function.
func printLines(days []string) {
	for _, day := range days {
		fmt.Println(day)
	}
}

func main() {
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday", "Saturday"}
	// YOUR CODE HERE: Get a slice containing just the
	// weekdays.
	weekdays := daysOfWeek[1:6]
	// Pass that slice to printLines.
	printLines(weekdays)
}