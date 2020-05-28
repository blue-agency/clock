# clock

This package provides a simple clock to get a current time.

## Usage

### sample

```go
package main

import (
    "github.com/kyohei-shimada/clock"
)

func main() {
    c := clock.New()
    c.Now() // Get a current time
}
```

on testing

```go
package main

import (
    "github.com/kyohei-shimada/clock/stub"
)

func Test_Someting(t *testing.T) {
    c := stub.New(time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC))
    c.Now() // 2009-11-10 23:00:00 +0000 UTC

    // ...

    c.Now() // get "2009-11-10 23:00:00 +0000 UTC" anytime
}
```
