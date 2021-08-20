package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func main() {
	lines, err := getStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	// YOUR CODE HERE: Build a slice containing all the
	// key strings from the "counts" map.
	var names []string
	for name := range counts {
		names = append(names, name)
	}

	// Call the sort.Strings method on your slice.
	sort.Strings(names)

	// Loop through the names in the sorted slice, and
	// print the name and the corresponding count from
	// the map.
	for _, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}
}