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
)

func main() {
	fmt.Print("Enter a grade: ") // Prompts the user to enter a grade
	reader := bufio.NewReader(os.Stdin) // Sets up a "buffered reader" that gets text from the keyboard
	input, err := reader.ReadString('\n') // Return everything the user has typed, up to where they pressed the Enter key.
	if err != nil { // If "error" is not nil
		log.Fatal(err) // Report the error and stop the program.
	}
	fmt.Println(input) // Print what the user typed.
}