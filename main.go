package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/thinhnn15/simple_bank/api"
	db "github.com/thinhnn15/simple_bank/db/sqlc"
	"github.com/thinhnn15/simple_bank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("cannot start server:", err)
	}
}
