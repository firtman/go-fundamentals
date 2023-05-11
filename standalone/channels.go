package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to send and receive messages.
	messages := make(chan string)

	// Go to a new goroutine and send a message.
	go func() {
		time.Sleep(2 * time.Second)
		messages <- "Hello, world!"
	}()

	// Wait for the message to be received.
	msg := <-messages

	// Print the message.
	fmt.Println(msg)
}
