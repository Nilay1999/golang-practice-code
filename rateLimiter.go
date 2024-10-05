package main

import (
	"fmt"
	"time"
)

const (
	MAX_BUCKET_SIZE = 3 // maximum capacity of the bucket
	REFILL_RATE     = 1 // refill rate in seconds
)

// Function to check if the bucket (queue) is full
func isBucketFull(queue []int) bool {
	return len(queue) >= MAX_BUCKET_SIZE
}

// Function to process incoming requests
func processRequest(queue *[]int, id int) {
	if isBucketFull(*queue) {
		fmt.Println("Request rejected! Bucket is full.")
		return
	}
	*queue = append(*queue, id) // add request to the queue
	fmt.Printf("Request accepted from %d. Queue: %v\n", id, *queue)
}

// Function to simulate leaking the bucket (processing requests)
func leakBucket(queue *[]int) {
	for {
		time.Sleep(REFILL_RATE * time.Second) // simulate time interval
		if len(*queue) > 0 {
			// Process the first request in the queue (leak)
			fmt.Printf("Processing request: %d. Queue before: %v\n", (*queue)[0], *queue)
			*queue = (*queue)[1:] // remove the processed request
			fmt.Printf("Queue after: %v\n", *queue)
		} else {
			fmt.Println("Bucket is empty, waiting for requests...")
		}
	}
}

func main() {
	queue := []int{}      // empty bucket
	go leakBucket(&queue) // start leaking requests in the background

	// Simulate incoming requests
	processRequest(&queue, 1)
	processRequest(&queue, 2)
	processRequest(&queue, 3)
	processRequest(&queue, 4) // This will be rejected since bucket is full
	processRequest(&queue, 5) // This will be rejected as well

	time.Sleep(10 * time.Second) // Allow some time for processing
}
