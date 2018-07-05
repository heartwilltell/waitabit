// Package waitabit gracefully handle shutdown for your application
package waitabit

import (
	"os"
	"os/signal"
	"sync"
	"time"
)

// Wait struct represent wait functionality
type Wait struct {
	wg      *sync.WaitGroup
	signals chan os.Signal
	events  chan struct{}
}

// NewWait create new example of Wait
func NewWait(signals ...os.Signal) (wait *Wait) {
	wait = &Wait{
		wg:      new(sync.WaitGroup),
		signals: make(chan os.Signal, 1),
		events:  make(chan struct{}),
	}
	signal.Notify(wait.signals, signals...)
	wait.wg.Add(1)
	wait.listen()
	return wait
}

// Wait simply wait for interruption
func (w *Wait) Wait() {
	w.wg.Wait()
}

// WaitWithFunc receive function as argument and call Wait() to wait interrupt signal. The function will be called after Wait()
func (w *Wait) WaitWithFunc(function func()) {
	w.wg.Wait()
	function()
}

// WaitWithTimeout wait certain amount fo time and then exit
func (w *Wait) WaitWithTimeout(timeout time.Duration) {
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		w.events <- struct{}{}
	}()
	w.wg.Wait()
}

// WaitWithTimeoutAndFunc wait certain amount of call func and then exit
func (w *Wait) WaitWithTimeoutAndFunc(timeout time.Duration, function func()) {
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		w.events <- struct{}{}
	}()
	w.wg.Wait()
	function()
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
