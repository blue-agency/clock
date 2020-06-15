# clock

This package provides a simple clock to get a current time.

## Usage

### sample

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/kyohei-shimada/clock"
)

func main() {
    // Get a current time as UTC
    c1, _ := clock.New(time.UTC)
    fmt.Printf("%+v\n", c1.Now()) // get a current time such as "2009-11-10 23:00:00 +0000 UTC"

	// Get a current time as another location
	jst, _ := time.LoadLocation("Asia/Tokyo")
	c2, _ := clock.New(jst)
	fmt.Printf("%+v\n", c2.Now()) // get a current time such as "2009-11-11 08:00:00 +0900 JST"
}
```

on testing

```go
package main

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/kyohei-shimada/clock/stub"
)

func TestSomething(t *testing.T) {
	c := stub.New(time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC))
	
    fmt.Printf("%+v\n", c.Now()) // get "2020-12-01 00:00:00 +0000 UTC"

	// spend several time
    time.Sleep(3 * time.Second)
	
	fmt.Printf("%+v\n", c.Now()) // get "2020-12-01 00:00:00 +0000 UTC" anytime
}
```
