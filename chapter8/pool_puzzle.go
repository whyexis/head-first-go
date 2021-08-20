/*
package geo

type Coordinates struct {
	Latitude  float64
	Longitude float64
*/

package main

import (
	"fmt"
	"geo"
)

func main() {
	location := geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
