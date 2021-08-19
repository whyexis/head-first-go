/*
 * Download and install the package with `go get github.com/jaymcgavren/car`
 */

package main

import (
	"github.com/jaymcgavren/car"
	"github.com/jaymcgavren/car/headlights"
	// All the .go source files in a package directory are
	// combined into a single package. So this one line
	// will give you access to the functions from both the
	// amplifier.go and speakers.go source files.
	"github.com/jaymcgavren/car/stereo"
	"github.com/jaymcgavren/car/wheels"
)
	
func main() {
	// Precede all function names with the package name.
	car.openDoor()
	// The package name is usually the same as the last
	// segment of the import path. So if you import the
	// "car/headlights" package, you'll refer to it as just
	// "headlights".
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()
}