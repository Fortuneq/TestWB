package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL, nats.Name("Friend API"),nats.DiscoveredServersHandler(func(c *nats.Conn) {
		fmt.Printf("Discovered servers %v", c.DiscoveredServers())
	}))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server connected")
	defer nc.Close()

}
