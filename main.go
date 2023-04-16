package main

import (
	"database/sql"
	"log"

	"github.com/ArsalanKm/simple_bank/api"
	db "github.com/ArsalanKm/simple_bank/db/sqlc"
	"github.com/ArsalanKm/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	config, err := util.LoadCofig(".")
	if err != nil {
		log.Fatal("cannot load configs")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect tot the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
