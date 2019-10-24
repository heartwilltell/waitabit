package waitabit

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func ExampleWait_Wait() {
	go func() {
		time.Sleep(500 * time.Millisecond)
		syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	wait := NewWait(os.Interrupt)
	wait.Wait()
	fmt.Println("Bye")
	// Output: Bye
}

func ExampleWait_WaitWithFunc() {
	go func() {
		time.Sleep(500 * time.Millisecond)
		syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	wait := NewWait(os.Interrupt)
	wait.WaitWithFunc(func() {
		// your logic here
		fmt.Println("Bye")
	})
	// Output: Bye
}

func ExampleWait_WaitWithFuncErr() {
	go func() {
		time.Sleep(500 * time.Millisecond)
		syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	wait := NewWait(os.Interrupt)
	if err := wait.WaitWithFuncErr(func() error {
		// your logic here
		return fmt.Errorf("something went wrong")
	}); err != nil {
		fmt.Println("something bad happened:", err)
	}
	fmt.Println("Bye")
	// Output: Bye
}

func ExampleWait_WaitWithTimeout() {
	go func() {
		time.Sleep(500 * time.Millisecond)
		syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	wait := NewWait(os.Interrupt)
	wait.WaitWithTimeout(1 * time.Second)
	fmt.Println("Bye")
	// Output: Bye
}

func ExampleWait_WaitWithTimeoutAndFunc() {
	go func() {
		time.Sleep(500 * time.Millisecond)
		syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	wait := NewWait(os.Interrupt)
	wait.WaitWithTimeoutAndFunc(1*time.Second, func() {
		// your logic here
		fmt.Println("Bye")
	})
	// Output: Bye
}

func ExampleWait_WaitWithTimeoutAndFuncErr() {
	go func() {
		time.Sleep(500 * time.Millisecond)
		syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	wait := NewWait(os.Interrupt)
	if err := wait.WaitWithTimeoutAndFuncErr(1*time.Second, func() error {
		// your logic here
		return fmt.Errorf("something went wrong")
	}); err != nil {
		fmt.Println("something bad happened:", err)
	}
	fmt.Println("Bye")
	// Output: Bye
}
