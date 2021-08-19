# Packages

* Keep related code together in one place
* Share code between programs or with other developers

## Go Workspace

* *bin*, holds compiled binary executable programs
* *pkg*, holds compiled binary package files
* *src*, holds Go source code

Within *src*, code for each package lives in its own separate subdirectory. By convention, the subdirectory name should be the same as the package name.

Each package directory should contain one or more source code files that end in a *.go* extension.

```
user's home directory
└───go (workspace directory)
    │───bin (holds executable programs)
    │───pkg (holds compiled package code)
    └───src (holds source code)
        │───doodad (package directory)
        └───gizmo
            │   gizmo.go
            │   plug.go
```

## Creating a new package

```go
package greeting // Package isn't main

import "fmt"

// Function names' first letters are capitalised so that functions are exported
func Hello() {
        fmt.Println("Hello!")
}

func Hi() {
        fmt.Println("Hi!")
}
```

Generally, the package name should match the name of the directory it's kept in, but the *main* package is an exception.