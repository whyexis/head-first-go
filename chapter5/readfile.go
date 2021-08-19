package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the data file for reading
	file, err := os.Open("data.txt") 
	if err != nil {
		log.Fatal(err)
	}

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file) 

	// Loops until the end of the file is reached
	// and scannerScan returns false
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Close the file to free resources
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}