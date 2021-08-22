package main

import "fmt"

type Fire struct {
	lit bool
}

func (f Fire) Light() {
	f.lit = true
	fmt.Println("Fire lit:", f.lit)
}
func (f Fire) Extinguish() {
	f.lit = false
	fmt.Println("Fire lit:", f.lit)
}

func Camp() error {
	var fire Fire
	fire.Light()
	defer fire.Extinguish()
	return fmt.Errorf("Error: spotted a bear")
	fmt.Println("Toasting marshmallows")
	return nil
}

func main() {
	err := Camp()
	if err != nil {
		fmt.Println(err)
	}
}