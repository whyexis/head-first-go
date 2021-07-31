/**
 * ------------------------------------------------------------------------
 * Exercise 3: for Loops
 * Here we have a program that’s meant to ask the user for a series of numbers and add them together.
 * Your goal is to modify the program so that it keeps asking the user for as many additional numbers as they want.
 * It should stop only when the user types “n” in response to the “Add more?” prompt.
 * (Notice that the “Y” in “Y/n” is capitalized, indicating that it’s the default. If the user hits Enter without typing anything, or indeed if they enter anything at all besides “n”, the program should ask for another number to add.)
 * Hint: You’ll probably want to use a for loop with no init or post statements to solve this.
 * For your first attempt, just make the program stop if the user enters “n”. But then see if you can make the program stop if the user enters “n” or “N”!
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
	 reader := bufio.NewReader(os.Stdin)
 
	 var sum float64
	 more := "y"

	 for {
		fmt.Print("Enter a number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}
	
		sum += number
	
		fmt.Print("Add more? [Y/n] ")
		more, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		more = strings.TrimSpace(more)
		more = strings.ToLower(more)

		if more == "n" {
			break
		}
	 }
 
	 fmt.Println("Sum is", sum)
 }