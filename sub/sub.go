package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main(){

	nc,err := nats.Connect(nats.DefaultURL,nats.Name("friend"))
	if err != nil{
		log.Fatal(err)
	}

	defer nc.Close()

	if err := nc.Publish("time",[]byte("How are your pecnic?"));err!=nil{
		log.Fatal(err)
	}
}