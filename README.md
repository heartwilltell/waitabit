# Wait a bit

<p>Tiny library for manage you application shutdown in a graceful way</p>

## Installation

```bash
go get -u github.com/heartwilltell/waitabit
```

## Usage

```go
package main

import (
    "os"
    "log"
)

import "github.com/heartwilltell/waitabit"


func main() { 
	
    // call your application here ...
    
    wait := waitabit.NewWait(os.Interrupt)
    wait.WaitWithFunc(func() {
        log.Println("Bye")
    })
}
```
