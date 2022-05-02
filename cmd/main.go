package main

import (
	"github.com/yura30rus/monopoly"
	"github.com/yura30rus/monopoly/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(monopoly.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
