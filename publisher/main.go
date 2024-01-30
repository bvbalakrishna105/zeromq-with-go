package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()

	// Bind to a specific address
	publisher.Bind("tcp://*:5555")

	for i := 0; ; i++ {
		// Send a message
		message := fmt.Sprintf("Message %d", i)
		publisher.Send(message, 0)
		time.Sleep(time.Second)
	}
}
