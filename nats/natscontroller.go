package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

func main() {
	var u User

	msg := natscontrol()
	databaseconnect(u,msg)

}

// Создал подключение к каналу
func natscontrol()(massege string){
	sc, err := stan.Connect("test-cluster", "friendlyAPi")

	if err != nil {
		fmt.Println(err)
	}

	if err = sc.Publish("nats-chan", []byte(`{"id":5,"user_name":"IT WORKS"}`)); err != nil {
		log.Fatal(err)
	}
	// Создал подписку на последнее сообщение в канале
	sub, err := sc.Subscribe("nats-chan", func(m *stan.Msg) {
		fmt.Printf("Recieved a message %s\n", string(m.Data))
		massege = string(m.Data)
	}, stan.StartWithLastReceived())
	if err != nil {
		fmt.Println(err)
	}

	defer sub.Unsubscribe()
	defer sc.Close()
	fmt.Println("End of nats")
	return 
}

type User struct{
	Id int `json:"id"`
	User_name string `json:"user_name"`
}

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "537j04222"
  dbname   = "postgres"
)

func databaseconnect(u User,message string){
	fmt.Println("Start of func")
	err := json.Unmarshal([]byte(message),&u);
	if err != nil {
		log.Print(err)
	}
	psqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s",host,port,user,password,dbname)
	db,err := sql.Open("postgres",psqlInfo)
	if err != nil{
		log.Print(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Successfully connected")
	_,gr := db.Exec(`INSERT INTO users(user_id,username) VALUES($1,$2)`, u.Id, u.User_name) 
	if gr != nil{
		log.Println(err)
	}

	fmt.Println(message)
}