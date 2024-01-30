package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()

	// Connect to the publisher's address
	subscriber.Connect("tcp://192.168.49.1:5555")

	// Subscribe to all messages
	subscriber.SetSubscribe("")

	for {
		// Receive a message
		message, _ := subscriber.Recv(0)
		fmt.Println("Received:", message)
	}
}
