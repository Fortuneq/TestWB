package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main(){
// Connect to a server
nc, err := nats.Connect(nats.DefaultURL, nats.Name("Friend API"))
if err != nil{
	log.Fatal(err)
}
getStatustxt:= func(nats.Conn)string{
	switch nc.Status(){
	case nats.CONNECTED:
		return "Connected"
	case nats.CLOSED:
		return "Closed"
	default:
		return "Other"
	}
}
fmt.Println(getStatustxt(*nc))
nc.Close()
fmt.Println(getStatustxt(*nc))
// Do smth..
}