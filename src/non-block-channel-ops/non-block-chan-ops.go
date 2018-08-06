package main

import (
	"fmt"
)

// Basic sends and receives on channels are blocking.
// However, we can use select with a default clause to implement non-blocking sends, receives,
// and even non-blocking multi-way selects.

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here’s a non-blocking receive. If a value is available on messages then select
	// will take the <-messages case with that value.
	// If not it will immediately take the default case
	select {
	case msg := <-messages:
		fmt.Println("Message Received: ", msg)
	default:
		fmt.Println("No Message Received")
	}

	// A non-blocking send works similarly. Here msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Message sent: ", msg)
	default:
		fmt.Println("No Message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals.
	select {
	case messages <- msg:
		fmt.Println("Message sent: ", msg)
	case sig := <-signals:
		fmt.Println("Signal Received: ", sig)
	default:
		fmt.Println("No Activity")

	}

}
