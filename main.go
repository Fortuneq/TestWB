package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main(){
// Connect to a server
nc, err := nats.Connect(nats.DefaultURL,nats.Name("Friend API"))
if err != nil{
	log.Fatal(err)
}
nc.Close()
}