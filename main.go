package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/thinhnn15/simple_bank/api"
	db "github.com/thinhnn15/simple_bank/db/sqlc"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@0.0.0.0:1701/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil{
		log.Fatal("cannot start server:", err)
	}
}
