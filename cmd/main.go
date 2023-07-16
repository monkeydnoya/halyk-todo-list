package main

import (
	"halyk-todo-list-api/internal/config"
	"halyk-todo-list-api/internal/controller/fiber"
	"halyk-todo-list-api/internal/service/todolist"
	conf "halyk-todo-list-api/pkg/config"
)

func main() {
	var config config.Config
	conf.Loader(&config)

	service := todolist.NewService(config)

	server := fiber.NewServer(config.Api, service)
	server.StartServer()
}
