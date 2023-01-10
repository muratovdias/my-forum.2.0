package main

import (
	"fmt"
	"log"

	"github.com/muratovdias/my-forum.2.0/config"
	"github.com/muratovdias/my-forum.2.0/internal/handler"
	"github.com/muratovdias/my-forum.2.0/internal/repository"
	"github.com/muratovdias/my-forum.2.0/internal/server"
	"github.com/muratovdias/my-forum.2.0/internal/service"
)

const port = ":8888"

func main() {
	configDB := config.NewConfDB()
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
