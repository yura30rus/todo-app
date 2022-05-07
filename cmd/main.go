package main

import (
	todo_app "github.com/yura30rus/todo-app"
	"github.com/yura30rus/todo-app/pkg/handler"
	"github.com/yura30rus/todo-app/pkg/repository"
	"github.com/yura30rus/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handdlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run("8000", handdlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
