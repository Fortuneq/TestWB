package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL, nats.Name("Friend API"), nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
		fmt.Printf("Client Disconnected %v", err)
	}), nats.ReconnectHandler(func(c *nats.Conn) { fmt.Println("Client Reconnect") }), nats.ClosedHandler(func(c *nats.Conn) {
		fmt.Println("Client Closed")
	}))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

}
