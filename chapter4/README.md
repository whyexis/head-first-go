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

## Package Naming Conventions

* All lowercase
* The name should be abbreviated if the meaning is fairly obvious
* One word, if possible. If two words are needed, they should not be separated by underscores, and the second word should not be capitalised.
* Don't use a name that package users are likely to use as well.

## Constants

Named values that never change

* `const` keyword
* Must assign a value at the time the constant is declared
* No short variable declaration syntax `:=`
* Type can be inferred from the value assigned
* Replace "hardcoded" literal values with constants, but name them with meaning
* Typically declared at the package level
* Constants that start with a capital letter are exported

## go install

Compiles and saves the executable in a standard directory

## Publishing packages

Use domain names and paths to ensure a package import path is unique e.g. `github.com/headfirstgo/keyboard`

## go get

Download and install packages. Sets up the subdirectories needed to set up the appropriate import path.