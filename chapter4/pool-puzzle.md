# Pool Puzzle

## Directory Set Up

```
user's home directory
└───go
    └───src
        └───calc
            │   calc.go
```

## Package Declaration

```go
package calc

func Add(first float64, second float64) float64 {
        return first + second
}

func Substract(first float64, second float64) float64 {
        return first - second
}
```

## Executable

```go
package main

import {
        "calc"
        "fmt"
}

func main() {
        fmt.Println(calc.Add(1,2))
        fmt.Println(calc.Substract(7,3))
}
```
