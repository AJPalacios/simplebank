package main

import (
	"database/sql"
	"log"

	"github.com/devspace/simplebank/api"
	db "github.com/devspace/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://myuser:mysecretpassword@localhost:5432/simple_bank?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
