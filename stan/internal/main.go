package main

import (
	"fmt"
	"log"
	"time"

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
	timer,_ := time.ParseDuration("30s")
	sub,err := sc.Subscribe("bro",func(m *stan.Msg){fmt.Printf("Received a message: %s\n", string(m.Data))},stan.StartAtTimeDelta(timer))
	if err != nil{
		log.Fatal(sub)
	}
	sub.Unsubscribe()

	sc.Close()
}