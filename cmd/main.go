package main

import (
	"log"

	"github.com/PrSmitch/first-app"
)

func main() {
	srv := new(first.Server)
	if err := srv.Run("8888"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())

	}

}
