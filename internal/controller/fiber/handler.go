package fiber

import (
	"context"
	"halyk-todo-list-api/internal/domain/models"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		Get Task
//	@Description	Get Task By Id
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string	true	"Task Id"
//	@Success		200	{object}	models.Task
//	@Router			/api/todo-list/v1/tasks/{id} [get]
func (s *Server) GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	task, err := s.Service.GetTask(requestCtx, id)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(task)
}

//	@Summary		Get Tasks
//	@Description	Get Task provided by Filters
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			Task	query		models.Filter	true	"Search Filter"
//	@Success		200		{object}	[]models.Task
//	@Router			/api/todo-list/v1/tasks [get]
func (s *Server) GetTasks(c *fiber.Ctx) error {
	filter := models.Filter{}
	if err := c.QueryParser(&filter); err != nil {
		return ErrorHandler(c, err)
	}

	requestCtx := context.Context(c.Context())
	tasks, err := s.Service.GetTasks(requestCtx, filter)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(tasks)
}

//	@Summary		Create Task
//	@Description	Create New Task
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			Task	body		models.CreateTask	true	"Create Task"
//	@Success		200		{object}	models.Task
//	@Router			/api/todo-list/v1/tasks [post]
func (s *Server) CreateTask(c *fiber.Ctx) error {
	newTask := models.CreateTask{}
	if err := c.BodyParser(&newTask); err != nil {
		return ErrorHandler(c, err)
	}

	requestCtx := context.Context(c.Context())
	task, err := s.Service.CreateTask(requestCtx, newTask)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(task)
}

//	@Summary		Update Task
//	@Description	Update existing Task
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			id		path		string				true	"Task Id"
//	@Param			Task	body		models.UpdateTask	true	"Update Task"
//	@Success		200		{object}	nil
//	@Router			/api/todo-list/v1/tasks/{id} [put]
func (s *Server) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	updateTask := models.UpdateTask{}
	if err := c.BodyParser(&updateTask); err != nil {
		return ErrorHandler(c, err)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.UpdateTask(requestCtx, id, updateTask)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Complete Task
//	@Description	Complete Task By Id
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string	true	"Task Id"
//	@Success		200	{object}	nil
//	@Router			/api/todo-list/v1/tasks/{id}/complete [put]
func (s *Server) CompleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.CompleteTask(requestCtx, id)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Complete Check List Item
//	@Description	Complete Check List Item Id
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string	true	"Check List Id"
//	@Success		200	{object}	nil
//	@Router			/api/todo-list/v1/tasks/checklist/{id}/complete [put]
func (s *Server) CompleteCheckListItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.CompleteCheckListItem(requestCtx, id)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Add CheckList Item To Task
//	@Description	Add CheckList Item To Task
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			id		path		string					true	"Task Id"
//	@Param			Task	body		models.CreateCheckList	true	"Create Check List Task"
//	@Success		200		{object}	nil
//	@Router			/api/todo-list/v1/tasks/{id}/checklist/add [put]
func (s *Server) AddCheckListItemToTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	newCheckList := models.CreateCheckList{}
	if err := c.BodyParser(&newCheckList); err != nil {
		return ErrorHandler(c, err)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.AddCheckListItemToTask(requestCtx, id, newCheckList)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Add Tag To Task
//	@Description	Add Existing Tag To Task
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			taskId	path		string	true	"Task Id"
//	@Param			tagId	path		string	true	"Tag Id"
//	@Success		200		{object}	nil
//	@Router			/api/todo-list/v1/tasks/{taskId}/tags/add/{tagId} [put]
func (s *Server) AddTagToTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	if taskId == "" {
		return c.SendStatus(400)
	}

	tagId := c.Params("tagId")
	if tagId == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.AddTagToTask(requestCtx, taskId, tagId)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Delete CheckList Item From Task
//	@Description	Delete CheckList Item From Task
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			taskId		path		string	true	"Task Id"
//	@Param			checkListId	path		string	true	"CheckList Id"
//	@Success		200			{object}	nil
//	@Router			/api/todo-list/v1/tasks/{taskId}/checklist/delete/{checkListId} [delete]
func (s *Server) DeleteCheckListItemFromTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	if taskId == "" {
		return c.SendStatus(400)
	}

	checkListId := c.Params("checkListId")
	if checkListId == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.DeleteCheckListItemFromTask(requestCtx, taskId, checkListId)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Delete Tag From Task
//	@Description	Delete Tag From Task
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			taskId	path		string	true	"Task Id"
//	@Param			tagId	path		string	true	"Tag Id"
//	@Success		200		{object}	nil
//	@Router			/api/todo-list/v1/tasks/{taskId}/tags/delete/{tagId} [delete]
func (s *Server) DeleteTagFromTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	if taskId == "" {
		return c.SendStatus(400)
	}

	tagId := c.Params("tagId")
	if tagId == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.DeleteTagFromTask(requestCtx, taskId, tagId)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Delete Task
//	@Description	Delete Task
//	@Tags			Tasks
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string	true	"Task Id"
//	@Success		200	{object}	nil
//	@Router			/api/todo-list/v1/tasks/{id} [delete]
func (s *Server) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.DeleteTask(requestCtx, id)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Get Tag
//	@Description	Get Tag By Id
//	@Tags			Tags
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string	true	"Tag Id"
//	@Success		200	{object}	models.Tag
//	@Router			/api/todo-list/v1/tags/{id} [get]
func (s *Server) GetTag(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	tag, err := s.Service.GetTag(requestCtx, id)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(tag)
}

//	@Summary		Get Tags
//	@Description	Get Tags
//	@Tags			Tags
//	@Accept			application/json
//	@Produce		json
//	@Success		200	{object}	[]models.Tag
//	@Router			/api/todo-list/v1/tags [get]
func (s *Server) GetTags(c *fiber.Ctx) error {
	requestCtx := context.Context(c.Context())
	tags, err := s.Service.GetTags(requestCtx)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(tags)
}

//	@Summary		Create Tag
//	@Description	Create New Tag
//	@Tags			Tags
//	@Accept			application/json
//	@Produce		json
//	@Param			Tag	body		models.CreateTag	true	"Create Tag"
//	@Success		200	{object}	models.Tag
//	@Router			/api/todo-list/v1/tags [post]
func (s *Server) CreateTag(c *fiber.Ctx) error {
	newTag := models.CreateTag{}
	if err := c.BodyParser(&newTag); err != nil {
		return ErrorHandler(c, err)
	}

	requestCtx := context.Context(c.Context())
	tag, err := s.Service.CreateTag(requestCtx, newTag)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.Status(200).JSON(tag)
}

//	@Summary		Update Tag
//	@Description	Update existing Tag
//	@Tags			Tags
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string				true	"Tag Id"
//	@Param			Tag	body		models.UpdateTag	true	"Update Tag"
//	@Success		200	{object}	nil
//	@Router			/api/todo-list/v1/tags/{id} [put]
func (s *Server) UpdateTag(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	updateTag := models.UpdateTag{}
	if err := c.BodyParser(&updateTag); err != nil {
		return ErrorHandler(c, err)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.UpdateTag(requestCtx, id, updateTag)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}

//	@Summary		Delete Tag
//	@Description	Delete existing Tag
//	@Tags			Tags
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string	true	"Tag Id"
//	@Success		200	{object}	nil
//	@Router			/api/todo-list/v1/tags/{id} [delete]
func (s *Server) DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}

	requestCtx := context.Context(c.Context())
	err := s.Service.DeleteTag(requestCtx, id)
	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.SendStatus(200)
}
