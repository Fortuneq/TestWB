package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL, nats.Name("Nice API"))
	if err != nil{
		log.Println(err)
	}
	defer nc.Close()
	
	sub,err := nc.SubscribeSync("updates")
	if err != nil{
		log.Fatal(err)
	}

	msg,err := sub.NextMsg(10*time.Second)
	if err != nil{
		log.Fatal(err)
	}

fmt.Printf("%s",msg.Data)
}