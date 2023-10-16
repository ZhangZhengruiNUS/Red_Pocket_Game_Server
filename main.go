package main

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	handler "Red_Pocket_Game_Server/handlers"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Afterwards, these variables will be changed to environment variables or configuration files
const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := handler.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
