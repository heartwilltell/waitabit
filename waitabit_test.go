package waitabit

import (
	"os"
	"syscall"
	"testing"
	"time"
)

func TestWait_Wait(t *testing.T) {

	t.Run("Wait", func(t *testing.T) {
		wait := NewWait(syscall.SIGTERM)
		syscall.Kill(os.Getpid(), syscall.SIGTERM)
		var result int
		wait.WaitWithFunc(func() {
			result = 1
		})
		if result != 1 {
			t.Error("Result is not equal 1")
		}
	})

	t.Run("Wait with timeout and func", func(t *testing.T) {
		start := time.Now().Unix()
		var result int
		wait := NewWait(syscall.SIGTERM)
		syscall.Kill(os.Getegid(), syscall.SIGTERM)
		wait.WaitWithTimeoutAndFunc(time.Second, func() {
			result = 1
		})
		end := time.Now().Unix()
		if end-start != 1 {
			t.Error("Something went wrong")
		}
		if result != 1 {
			t.Error("Result is not equal 1")
		}
	})

	t.Run("Wait with func", func(t *testing.T) {
		wait := NewWait(syscall.SIGTERM)
		syscall.Kill(os.Getpid(), syscall.SIGTERM)
		wait.Wait()
	})

	t.Run("Wait with timeout", func(t *testing.T) {
		start := time.Now().Unix()
		wait := NewWait(os.Interrupt)
		wait.WaitWithTimeout(time.Second)
		end := time.Now().Unix()
		if end-start != 1 {
			t.Error("Something went wrong")
		}
	})
}
