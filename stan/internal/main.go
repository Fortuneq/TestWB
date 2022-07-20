package main

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

func main(){
	sc, err := stan.Connect( "test-cluster", "vlad",stan.MaxPubAcksInflight(1000))
	if err != nil{
		log.Fatal(err)
	}
	if err := sc.Publish("bro",[]byte(" World song"));err!=nil{
		log.Fatal(err)
	}
	log.Println("Connected")

	sub,err := sc.Subscribe("bro",func(msg *stan.Msg){msg.Ack()},stan.StartWithLastReceived(),stan.DurableName("durable-sub"))
	if err != nil{
		log.Fatal(sub)
	}
	fmt.Println(sub)
	sub.Close()

	sc.Close()
}