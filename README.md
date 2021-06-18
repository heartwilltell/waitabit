# Wait a bit

Tiny library for manage you application shutdown in graceful way by catching the OS signals.

**⚠️ This repository is now archived. Please use `os/signal.NotifyContext` instead**


## Documentation

[![](https://goreportcard.com/badge/github.com/heartwilltell/waitabit)](https://goreportcard.com/report/github.com/heartwilltell/waitabit)
[![](https://godoc.org/github.com/heartwilltell/waitabit?status.svg)](https://pkg.go.dev/github.com/heartwilltell/waitabit?tab=doc)

## Installation

```bash
go get -u github.com/heartwilltell/waitabit
```

## Usage

```go
package main

import (
    "log"
    "os"

    "github.com/heartwilltell/waitabit"
)


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
