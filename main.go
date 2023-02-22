package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MeganViga/Plipkart/api"
	db "github.com/MeganViga/Plipkart/db/sqlc"
	"github.com/MeganViga/Plipkart/util"
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
	config, err := util.LoadConfig(".")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	conn, err :=  sql.Open("postgres",psqlInfo)
	if err != nil{
		panic(err)
	}
	store := db.NewStore(conn)
	server , err := api.NewServer(config,store)
	if err != nil{
		log.Fatal(err)
	}
	server.StartServer(config.HTTPServerAddress)


}