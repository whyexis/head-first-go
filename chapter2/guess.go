/**
 * ------------------------------------------------------------------------
 * Guess Game.
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
	"fmt"
)

func main() {
	fmt.Println("Please guess the target number: ")
 }