package main

import (
	"log"
	"github.com/Lzrb0x/estiloMestreGO/internal/server"
	"github.com/Lzrb0x/estiloMestreGO/internal/db"
)




func main() {
	
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	
	server := server.NewServer(dbConnection)
	log.Printf("Starting server on %s\n", server.Addr)

	err = server.ListenAndServe()
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}