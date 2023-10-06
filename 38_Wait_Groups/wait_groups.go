package main

import (
	"fmt"
	"sync"
	"time"
)

/*
To wait for multiple goroutines to finish, we can use a wait group.
*/

// This is the function we'll run in every goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// Sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This Wait Group is used to wait for all the gouroutines launched
	// here to finish.
	// Note: if a WaitGroup is passed into functions,
	// it should be done by pointer
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Launch several goroutines
		// and increment the WaitGroup counter for each.
		wg.Add(1)

		// Avoid re-use of the same i value in each goroutine closure.
		// See the FAQ (https://golang.org/doc/faq#closures_and_goroutines)
		// for more details.
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they’re done.
	wg.Wait()

	// Note that this approach has no straightforward way to propagate errors
	// from workers.
	// For more advanced use cases, consider using the errgroup package
}
