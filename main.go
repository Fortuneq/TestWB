package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main(){
// Connect to a server
nc, err := nats.Connect(nats.DefaultURL, nats.Name("Friend API"), nats.Token("tokens"))
if err != nil{
	log.Fatal(err)
}
fmt.Println("Get connected")
 defer nc.Close()
// Do smth..
}