package main

import (
	"log"

	"github.com/Lzrb0x/estiloMestreGO/internal/db"
	"github.com/Lzrb0x/estiloMestreGO/internal/server"
)

func main() {

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	newServer := server.NewServer(dbConnection)
	log.Printf("Starting server on %s\n", newServer.Addr)

	err = newServer.ListenAndServe()
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
