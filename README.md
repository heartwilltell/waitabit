# Wait a bit

<p>Tiny library for manage you application shutdown in a graceful way</p>

## Installation

```bash
go get -u github.com/heartwilltell/waitabit
```

## Usage

```
package main


func main() {   
    
    // call your application here ...
    
    wait := NewWait(os.Interrupt)
    wait.WaitWithFunc(func() {
        logger.Debug("Bye")
    })
    
}
```
