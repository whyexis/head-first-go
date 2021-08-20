Package geo

```go
package geo

type Coordinates struct {
        Latitude float64
        Longitude float64
}

type Landmark struct {
        Name string
        Coordinates
}
```

Main

```go
package main

import (
        "fmt"
        "geo"
)

func main() {
        location := geo.Landmark()
        location.Name = "The Googleplex"
        location.Latitude = 37.42
        location.Longitude = -122.08
        fmt.Println(location)
}

```