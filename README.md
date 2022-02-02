# stack-via-golang

## Installing
First, use `go get` to install the latest version of `pattern` package.

This command will install the `pattern` the library and its dependencies:

    go get -u github.com/prasnitt/stack-via-golang/pattern

Next, include `pattern` in your application:

```go
import "github.com/prasnitt/stack-via-golang/pattern"
```

## Usage
The `Pattern` package provides function (`IsValid`) to check if a corresponding package is valid or not.

### Example
Example Code: 

```go
package main

import (
	"fmt"

	"github.com/prasnitt/stack-via-golang/pattern"
)

func main() {
	testPatterns := []string{
		"{}",
		"{[]}",
		"{[}",
		"{}()[{}(){}]",
		"{[]}()[{}(){}]",
		"{你好世界 (Hello World) A[B]C}(D)E[{F}(G){H}I]J",
	}

	for _, t := range testPatterns {

		if pattern.IsValid(t) {
			fmt.Printf(" Pattern '%s' is Valid\n", t)
		} else {
			fmt.Printf(" Pattern '%s' is Invalid\n", t)
		}
	}

}

```

You will see following output after running the above code
```
 Pattern '{}' is Valid
 Pattern '{[]}' is Valid
 Pattern '{[}' is Invalid
 Pattern '{}()[{}(){}]' is Valid
 Pattern '{[]}()[{}(){}]' is Valid
 Pattern '{你好世界 (Hello World) A[B]C}(D)E[{F}(G){H}I]J' is Valid
```

## Steps Taken to build this go module:

| # |  Short descript  | Status |
|:-----|:--------------:|------:|
| 0   |  Setup Git Repo | DONE |
| 1   |  Setup Go lang Environment (e.g. Go modules)  |   DONE |
| 2   | Create `stack` package (with Unit test) |   DONE | 
| 3   | Create `pattern` package (with Unit test)|   DONE | 
| 4   | Document usage |   DONE | 
| 5   | Final tidy up and review |   DONE | 
