package main

import "fmt"

var WorkerQueue chan chan WorkRequest

func StartDispatcher(nWorkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nWorkers)

	// Now, create all of our workers.
	for i := 0; i < nWorkers; i++ {
		fmt.Println("Starting worker", i + 1)
		worker := NewWorker(i + 1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
