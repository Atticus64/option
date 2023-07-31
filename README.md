# Option in Golang

* Option structure inspired in Rust :crab:

## Installation

```bash
go get github.com/atticus64/option
```

Example:

```go
package main

import (
	"fmt"
	"github.com/atticus64/option"
)

func main() {
	data := option.Optional[int]()

	if data.IsNone() {
		val := data.ValueOr(78)
		fmt.Println(val)
		fmt.Println("is none")
	} else {
		value, _ := data.GetValue()
		fmt.Println(value)
	}
}
```

