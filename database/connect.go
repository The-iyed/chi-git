package database

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
  )

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "golangdb"
  )

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)   
	if err != nil {
        log.Fatalln(err)
    }
  
    defer db.Close()
 
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully Connected")
    }
}