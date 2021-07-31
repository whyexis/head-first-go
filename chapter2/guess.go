/**
 * ------------------------------------------------------------------------
 * Guess challenges player to guess a random number.
 * Generate a random number from 1 to 100, and
 * store it as a target number for the player to guess.
 * Prompt the player to guess what the target number is,
 * and store their response.
 * If the guess is less than the target, say,
 * "Opps. Your guess was LOW.", and vice versa.
 * Allow the player to guess up to 10 times.
 * Before each guess, let them know how many guesses they have left.
 * If the player guessed the target number correctly,
 * tell them, "Good job! You guessed it!".
 * Then stop asking for new guesses.
 * If the player ran out of turns without guessing correctly, say,
 * "Sorry. You didn't guess my number. It was: [target]"
 * ------------------------------------------------------------------------
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand" // Full import path, package name: "rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // Get the current date and time as an integer
	rand.Seed(seconds) // Seed the random number generator
	target := rand.Intn(100) + 1 // Generates an integer from 0 to 99, add 1 to make it from 1 to 100
	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 10; guesses > 0; guesses -- {
		fmt.Println("You have", guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Opps. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Opps. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break 
		}
	}

	if !success {
		fmt.Println("Sorry. You didn't guess my number. It was:", target)
	}
 }