/**
 * ------------------------------------------------------------------------
 * Fill in the Blanks
 * The program below should take a number of cans of food and calculate their total weight.
 * It should declare a count variable with the number of cans, and a unitWeight variable with the amount (in kilograms) that each can weighs.
 * It should then multiply count by unitWeight, and assign the result to a new totalWeight variable.
 * Finally, it should print the number of cans and their total weight.
 * Fill in the blanks in the program so that it will run and produce the specified output. Note that count will have a type of int, but unitWeight will have a type of float64, so you’ll need to do a conversion before you can multiply the two together.
 * ------------------------------------------------------------------------
 */
 package main

 import "fmt"
 
 func main() {
     var count int = 20
     unitWeight := 0.4 // Short variable declarations do not need "var" at the beginning, will return an error
     totalWeight := float64(count) * unitWeight
     fmt.Println(count, "cans weigh", totalWeight, "kilograms")
 }