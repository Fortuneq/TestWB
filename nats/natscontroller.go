package main

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

func main(){
	// Создал подключение к каналу
	sc ,err := stan.Connect("test-cluster","friendlyAPi")
	
	if err != nil {
		fmt.Println(err)
	}

	if err = sc.Publish("nats-chan",[]byte("Hello bro"));err != nil{
		log.Fatal(err)
	}
	// Создал подписку на последнее сообщение в канале
	sub,err := sc.Subscribe("nats-chan",func(m *stan.Msg){fmt.Printf("Recieved a message %s\n",string(m.Data))},stan.StartWithLastReceived())
	if err != nil{
		fmt.Println(err)
	}

	
	defer sub.Unsubscribe()
	defer sc.Close()
	
}