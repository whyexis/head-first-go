package main

import "fmt"

func main() {
	weekdays := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for i, day := range weekdays {
		fmt.Println(i, day)
	}
}