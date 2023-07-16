package fiber

import (
	"halyk-todo-list-api/internal/config"
	"halyk-todo-list-api/internal/controller"
	"halyk-todo-list-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	ApiConfig config.Api
	Service   service.TodoListService
	App       *fiber.App
}

//	@title			Halyk Todo List API
//	@version		0.0.1
//	@description	Halyk Todo List API
//	@schemes		http
func NewServer(config config.Api, service service.TodoListService) controller.Controller {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
	})
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, X-Api-Key, X-Requested-With, Content-Type, Accept, Authorization, authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	server := Server{
		Service:   service,
		App:       app,
		ApiConfig: config,
	}

	server.SetupInternalRoutes()
	server.SetupRoutes()
	return &server
}
