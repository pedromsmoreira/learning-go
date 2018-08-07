package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization
// and maintaining quality of service.
// Go elegantly supports rate limiting with goroutines, channels, and tickers.

func main() {
	// First we’ll look at basic rate limiting.
	// Suppose we want to limit our handling of incoming requests.
	// We’ll serve these requests off a channel of the same name.

	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds.
	// This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request,
	// we limit ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("Request: ", req, time.Now().UTC())
	}

	// We may want to allow short bursts of requests in our rate limiting scheme
	// while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel.
	// This burstyLimiter channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now().UTC()
	}

	// Every 200 milliseconds we’ll try to add a new value to burstyLimiter,
	// up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests.
	// The first 3 of these will benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Request: ", req, time.Now().UTC())
	}

	/*

		Running our program we see the first batch of requests handled once every ~200 milliseconds
		as desired.

		$ go run rate-limiting.go
		Request:  0 2018-08-07 00:54:48.180865275 +0000 UTC
		Request:  1 2018-08-07 00:54:48.37815205 +0000 UTC
		Request:  2 2018-08-07 00:54:48.576904537 +0000 UTC
		Request:  3 2018-08-07 00:54:48.776876722 +0000 UTC
		Request:  4 2018-08-07 00:54:48.97633472 +0000 UTC

		For the second batch of requests we serve the first 3 immediately because
		of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.

		Request:  0 2018-08-07 00:54:48.97644069 +0000 UTC
		Request:  1 2018-08-07 00:54:48.9764542 +0000 UTC
		Request:  2 2018-08-07 00:54:48.976461597 +0000 UTC
		Request:  3 2018-08-07 00:54:49.176973666 +0000 UTC
		Request:  4 2018-08-07 00:54:49.376967068 +0000 UTC
	*/
}
