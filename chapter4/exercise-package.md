```
user's home directory
└───go
    └───src
        └───my.com
            └───me
                └───myproject
                    └───mypackage
                        └───mypackage.go

```

mypackage.go

```go
package mypackage

func MyFunction() {

}
```

Answer

```go
package main

import "my.com/me/myproject/mypackage"

func main () {
        mypackage.MyFunction()
}
```
