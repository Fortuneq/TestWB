package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "537j04222"
  dbname   = "postgres"
)

func main(){
	psqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s",host,port,user,password,dbname)
	db,err := sql.Open("postgres",psqlInfo)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Successfully connected")
}