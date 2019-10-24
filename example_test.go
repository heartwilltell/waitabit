package waitabit

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ExpampleWait() {
	wait := NewWait(os.Interrupt)
	wait.Wait()
	log.Println("Bye")
}

func ExpampleWaitWithFunc() {
	wait := NewWait(os.Interrupt)
	wait.WaitWithFunc(func() {
		// your logic here
		log.Println("Bye")
	})
}

func ExpampleWaitWithFuncErr() {
	wait := NewWait(os.Interrupt)
	if err := wait.WaitWithFuncErr(func() error {
		// your logic here
		return fmt.Errorf("something went wrong")
	}); err != nil {
		log.Println("something bad happened:", err)
	}
	log.Println("Bye")
}

func ExpampleWaitWithTimeout() {
	wait := NewWait(os.Interrupt)
	wait.WaitWithTimeout(1 * time.Second)
	log.Println("Bye")
}

func ExpampleWaitWithTimeoutAndFunc() {
	wait := NewWait(os.Interrupt)
	wait.WaitWithTimeoutAndFunc(1*time.Second, func() {
		// your logic here
		log.Println("Bye")
	})
}

func ExpampleWaitWithTimeoutAndFuncErr() {
	wait := NewWait(os.Interrupt)
	if err := wait.WaitWithTimeoutAndFuncErr(1*time.Second, func() error {
		// your logic here
		return fmt.Errorf("something went wrong")
	}); err != nil {
		log.Println("something bad happened:", err)
	}
	log.Println("Bye")
}
