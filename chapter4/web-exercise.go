/*
 * Download and install the package with `go get github.com/jaymcgavren/car`
 */

package main

import (
	"github.com/jaymcgavren/car"
	"github.com/jaymcgavren/car/headlights"
	"github.com/jaymcgavren/car/stereo"
	"github.com/jaymcgavren/car/wheels"
)
	
func main() {
	car.openDoor()
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()
}