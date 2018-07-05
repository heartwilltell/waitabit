// Package waitabit gracefully handle shutdown for your application
package waitabit

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Wait struct represent wait functionality.
type Wait struct {
	wg      *sync.WaitGroup
	signals chan os.Signal
	events  chan struct{}
}

// NewWait create a new instance of Wait type.
func NewWait(signals ...os.Signal) *Wait {
	wait := &Wait{
		wg:      &sync.WaitGroup{},
		signals: make(chan os.Signal, 1),
		events:  make(chan struct{}),
	}
	signal.Notify(wait.signals, signals...)
	wait.wg.Add(1)
	wait.listen()
	return wait
}

// Wait simply wait for interruption.
func (w *Wait) Wait() {
	w.wg.Wait()
}

// WaitWithFunc receive function as argument and call Wait() to wait interrupt signal.
// The function will be called after Wait().
func (w *Wait) WaitWithFunc(function func()) {
	w.wg.Wait()
	function()
}

// WaitWithFuncErr receive function as argument and call Wait() to wait interrupt signal.
// The function will be called after Wait().
// If function return an error it would be wrapped and returned.
func (w *Wait) WaitWithFuncErr(function func() error) error {
	w.wg.Wait()
	if err := function(); err != nil {
		return fmt.Errorf("function call returned an error: %w", err)
	}
	return nil
}

// WaitWithTimeout wait a given amount fo time and then exit.
func (w *Wait) WaitWithTimeout(timeout time.Duration) {
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		w.events <- struct{}{}
	}()
	w.wg.Wait()
}

// WaitWithTimeoutAndFunc wait a given amount of time then call function and then exit.
func (w *Wait) WaitWithTimeoutAndFunc(timeout time.Duration, function func()) {
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		w.events <- struct{}{}
	}()
	w.wg.Wait()
	function()
}

// WaitWithTimeoutAndFunc wait a given amount of call function and then exit.
// If function return an error it would be wrapped and returned.
func (w *Wait) WaitWithTimeoutAndFuncErr(timeout time.Duration, function func() error) error {
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		w.events <- struct{}{}
	}()
	w.wg.Wait()
	if err := function(); err != nil {
		return fmt.Errorf("function call returned an error: %w", err)
	}
	return nil
}

// listen a simple function that listen to the signals channel for interruption signals and then call Done() of the wait group
func (w *Wait) listen() {
	go func() {
		defer w.wg.Done()
		for {
			select {
			case <-w.signals:
				return
			case <-w.events:
				return
			}
		}
	}()
}
