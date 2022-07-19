package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main(){
// Connect to a server
nc, err := nats.Connect("vlad:qwerty@127.0.0.1")
if err != nil{
	log.Fatal(err)
}
fmt.Println("Get connected")
 defer nc.Close()
// Do smth..
}