package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {
	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters.\n", l, l.ToMilliliters())
	ml := Milliliters(500)
	fmt.Printf("%0.1f milliliters is %0.1f liters.\n", ml, ml.ToLiters())
}