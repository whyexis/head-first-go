coordinates.go

```go
package geo

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func (c *Coordinates) SetLatitude(latitude float64) {
	c.Latitude = latitude
}

func (c *Coordinates) SetLongitude(longitude float64) {
	c.Longitude = longitude
}
```

main.go

```go
package main

import (
	"fmt"
	"geo"
)

func main() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
```
