package main

import (
	"os"
	"payment-service/db"
	"payment-service/internal/config"
	"payment-service/internal/server"
)

func main() {
	config.Init()

	database := db.Database()

	app := server.NewApp(database)

	port := os.Getenv("GRPC_PORT")
	go config.StartGRPCServer(database, port)

	app.RunHTTP(config.GetPort())
}
