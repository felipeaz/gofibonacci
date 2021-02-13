# gofibonacci
Go Fibonacci is a simple package that generates a Fibonacci Sequence with the given amount of numbers and also returns a number from the given position.

# Installation
Use `go tool` to install this package on your go Project.

`go get -u github.com/FelipeAz/gofibonacci`

# Usage
You must have this package installed on your project. You can do that by using the line bellow.
```
import (
	"fmt"

	fibonacci "github.com/FelipeAz/gofibonacci"
)
```

# Example
```
package main

import (
	"fmt"

	fibonacci "github.com/FelipeAz/gofibonacci"
)

func main() {
	fmt.Println(fibonacci.Sequence(20))
	fmt.Println(fibonacci.NumberAtPosition(10))
}
```

## License
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat)](LICENSE)
