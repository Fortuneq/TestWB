package main
import "github.com/nats-io/nats.go"

func main(){
// Connect to a server
nc, _ := nats.Connect(nats.DefaultURL)

nc.Close()
}