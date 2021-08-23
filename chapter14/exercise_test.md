math.go
```go
package arithmetic

func Add(a float64, b float64) float64 {
        return a + b
}

func Subtract(a float64, b float64) float64 {
        return a - b
}
```
math_test.go
```go
package arithmetic

import "testing"

func TestAdd(t *testing.T) {
        if Add(1, 2) != 3 {
                t.Error("1 + 2 did not equal 3")
        }
}

func TestSubtract(t *testing.T) {
        if Subtract(8, 4) != 4 {
                t.Error("8 -4 did not equal 4")
        }
}
```