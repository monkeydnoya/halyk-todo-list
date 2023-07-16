package fiber

import (
	_ "halyk-todo-list-api/internal/controller/fiber/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (s *Server) SetupInternalRoutes() {
	internalRoutes := s.App.Group("/internal")
	internalRoutes.Get("/metrics", monitor.New(monitor.Config{Title: "Halyk TodoList Monitoring Page"}))

	internalRoutes.Get("/docs/*", swagger.New())
}

func (s *Server) SetupRoutes() {
	routes := s.App.Group("/api/todo-list")
	v1 := routes.Group("/v1")
	v1.Get("/tasks", s.GetTasks)
	v1.Get("/tasks/:id", s.GetTask)
	v1.Post("/tasks", s.CreateTask)
	v1.Delete("/tasks/:id", s.DeleteTask)

	v1.Put("/tasks/:id", s.UpdateTask)
	v1.Put("/tasks/:id/complete", s.CompleteTask)

	v1.Put("/tasks/checklist/:id/complete", s.CompleteCheckListItem)
	v1.Put("/tasks/:id/checklist/add", s.AddCheckListItemToTask)
	v1.Put("/tasks/:taskId/tags/add/:tagId", s.AddTagToTask)

	v1.Delete("/tasks/:taskId/checklist/delete/:checkListId", s.DeleteCheckListItemFromTask)
	v1.Delete("/tasks/:taskId/tags/delete/:tagId", s.DeleteTagFromTask)
	v1.Delete("/tasks/:taskId/tags/delete/:tagId", s.DeleteTagFromTask)

	v1.Get("/tags", s.GetTags)
	v1.Get("/tags/:id", s.GetTag)
	v1.Post("/tags", s.CreateTag)
	v1.Put("/tags/:id", s.UpdateTag)
	v1.Delete("/tags/:id", s.DeleteTag)
}
