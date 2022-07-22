package main

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL, nats.Name("friend"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	type time struct {
		Time string
		What string
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := ec.Subscribe("time", func(t *time) {
		log.Printf("Your t is %s and now %s", t.Time, t.What)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
