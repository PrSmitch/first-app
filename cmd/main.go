package main

import (
	"log"

	"github.com/PrSmitch/first-app"
	handler "github.com/PrSmitch/first-app/pkg/handlers"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(first.Server)
	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())

	}

}
