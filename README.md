# Wait a bit

<p>Tiny library for manage you application shutdown in a graceful way</p>

## Documentation

[![](https://goreportcard.com/badge/github.com/heartwilltell/waitabit)](https://goreportcard.com/report/github.com/heartwilltell/waitabit)
[![](https://godoc.org/github.com/heartwilltell/waitabit?status.svg)](https://godoc.org/github.com/heartwilltell/waitabit)

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
    
    // or ...
    
    waitabit.NewWait(os.Interrupt).WaitWithFunc(func() {
    	log.Println("Bye")
    })
    
}
```
