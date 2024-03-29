/**
 * ------------------------------------------------------------------------
 * Pass Fail
 * Reports whether a grade passes or fails.
 * ------------------------------------------------------------------------
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ") // Prompts the user to enter a grade
	reader := bufio.NewReader(os.Stdin) // Sets up a "buffered reader" that gets text from the keyboard
	input, err := reader.ReadString('\n') // Return everything the user has typed, up to where they pressed the Enter key.
	if err != nil { // If "error" is not nil
		log.Fatal(err) // Report the error and stop the program.
	}
	
	input = strings.TrimSpace(input) // Trim the newline character from the input string
	grade, err := strconv.ParseFloat(input, 64) // Convert the string to a float64 value.
	if err != nil {
		log.Fatal(err)
	}

	var status string // Declare the "status" variable here, so it is in scope for the rest of the function
	if grade >= 60 { // Compare to the float64 in "grade" not the string in input
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status) // Print the entered grade and the pass/fail status.
}