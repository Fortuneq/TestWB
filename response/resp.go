package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL, nats.Name("Friend"))
	e(err)

	defer nc.Close()

	sub,err := nc.SubscribeSync("time")
	e(err)

	msg,err := sub.NextMsg(10 * time.Second)

	e(err)
	timeAsBytes := []byte(time.Now().String())

	msg.Respond(timeAsBytes)
}

func e(err error){
	if err != nil{
		log.Fatal(err)
	}
}