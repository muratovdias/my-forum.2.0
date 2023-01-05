package main

import (
	"fmt"
	"forum/internal/handler"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	"log"
)

const port = ":8888"

func main() {
	configDB := repository.NewConfDB()
	db, err := repository.InitDB(configDB)
	if err != nil {
		log.Fatalf("failed to initialize db : %s", err.Error())
	}
	if err := repository.CreateTables(db); err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	services := service.NewService(*repo)
	handler := handler.NewHandler(services)

	server := new(server.Server)
	fmt.Printf("Starting server at port %s\nhttp://localhost%s/\n", port, port)
	if err := server.Run(port, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
