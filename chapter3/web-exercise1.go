/**
 * ------------------------------------------------------------------------
 * Student Tests
 * Write a scoreSummary function that accepts a string parameter with the student name, followed by three float64 parameters with the test scores. Calculate the average of the scores by adding them all together and dividing the result by three. Then call fmt.Printf to print a table row with the following columns:
 *  The student name string, with a width of 10 characters.
 *  The first test score, with a with of 8 characters, 2 of them decimal places.
 *  The second test score, with a with of 8 characters, 2 of them decimal places.
 *  The third test score, with a with of 8 characters, 2 of them decimal places.
 *  The average of the scores, with a with of 8 characters, 2 of them decimal places.
 * Separate the columns with a vertical bar surrounded by spaces, just like you see in the call to Printf in the main function.
 * ------------------------------------------------------------------------
 */

package main

import "fmt"

// scoreSummary takes four parameters.
// "name": "string", "score1", "score2", "score3": "float64".
func scoreSummary(name string, score1 float64, score2 float64, score3 float64) {
	// The parentheses ensure that the addition operations
	// take place before dividing.
	average := (score1 + score2 + score3) / 3
	// %s formats a string value.
	// %10s pads a string with spaces to a width of 10.
	// %f formats a floating-point value.
	// %8.2f pads a floating-point value with spaces to a
	// width of 8, ensuring that a minimum of two decimal
	// places are displayed.
	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n", name, score1, score2, score3, average)
}

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println("------------------------------------------------------")
	scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	scoreSummary("Jessie", 79.3, 99.1, 82.5)
	scoreSummary("Lamar", 82.2, 95.4, 77.6)
}