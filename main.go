package main

import (
	"database/sql"
	"fmt"

	"github.com/MeganViga/Plipkart/api"
	_ "github.com/lib/pq"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "plipkart"
  )
func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err :=  sql.Open("postgres",psqlInfo)
	if err != nil{
		panic(err)
	}
	server := api.NewServer(db)
	server.StartServer()


}