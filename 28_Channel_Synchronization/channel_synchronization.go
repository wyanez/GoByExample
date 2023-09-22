package main

import (
	"fmt"
	"time"
)

/*
We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
*/

// This is the function we’ll run in a goroutine.
// The done channel will be used to notify another goroutine
// that this function’s work is done.
func worker(done chan bool) {
	fmt.Println("[worker] working...")
	time.Sleep(time.Second)

	// Send a value to notify that we’re done.
	done <- true
	fmt.Println("[worker] done")
}

func main() {
	done := make(chan bool, 1)
	//Start a worker goroutine, giving it the channel to notify on.
	go worker(done)

	//Block until we receive a notification from the worker on the channel.
	fmt.Println("[main] waiting")
	<-done
	// IMPORTANT! If you removed the <- done line from this program,
	// the program would exit before the worker even started.
	fmt.Println("[main] done!")
}
