package postgres

import (
	"context"
	"halyk-todo-list-api/internal/data/database/postgres/entities"
	"halyk-todo-list-api/internal/domain/models"
	"strconv"
	"time"
)

func (pg *pgConnection) GetTask(ctx context.Context, id string) (models.Task, error) {
	var taskEntity entities.Task

	query := pg.db.WithContext(ctx).Model(&entities.Task{})
	if err := query.Where("id = ?", id).First(&taskEntity).Error; err != nil {
		pg.log.Errorw("postgres: could not found", err)
		return models.Task{}, pg.convertError(err)
	}

	checkList := make([]entities.CheckList, 0)
	tags := make([]entities.Tag, 0)
	pg.db.WithContext(ctx).Model(&taskEntity).Association("CheckList").Find(&checkList)
	pg.db.WithContext(ctx).Model(&taskEntity).Association("Tags").Find(&tags)

	taskEntity.CheckList = checkList
	taskEntity.Tags = tags

	return TaskEntityToModel(taskEntity), nil
}

func (pg *pgConnection) GetCheckListItem(ctx context.Context, id string) (models.CheckList, error) {
	var checkListEntity entities.CheckList

	query := pg.db.WithContext(ctx).Model(&entities.CheckList{})
	if err := query.Where("id = ?", id).First(&checkListEntity).Error; err != nil {
		pg.log.Errorw("postgres: could not found", err)
		return models.CheckList{}, pg.convertError(err)
	}

	return CheckListEntityToModel(checkListEntity), nil
}

func (pg *pgConnection) GetTasks(ctx context.Context, filter models.Filter) ([]models.Task, error) {
	var tasksEntities []entities.Task

	query := pg.db.WithContext(ctx).Model(&entities.Task{})
	if filter.Done != "" {
		done, err := strconv.ParseBool(filter.Done)
		if err != nil {
			pg.log.Errorw("postgres: done is wrong type")
			return []models.Task{}, pg.convertError(err)
		}
		query = query.Where("done = ?", done)
	}
	if filter.CreatedFrom != 0 {
		query = query.Where("created_at >= ?", filter.CreatedFrom)
	}
	if filter.CreatedTo != 0 {
		query = query.Where("created_at <= ?", filter.CreatedTo)
	}

	if err := query.Find(&tasksEntities).Error; err != nil {
		pg.log.Errorw("postgres: could not found", err)
		return []models.Task{}, pg.convertError(err)
	}

	for _, taskEntity := range tasksEntities {
		checkList := make([]entities.CheckList, 0)
		tags := make([]entities.Tag, 0)
		pg.db.WithContext(ctx).Model(&taskEntity).Association("CheckList").Find(&checkList)
		pg.db.WithContext(ctx).Model(&taskEntity).Association("Tags").Find(&tags)

		taskEntity.CheckList = checkList
		taskEntity.Tags = tags
	}

	var tasks []models.Task
	for _, taskEntity := range tasksEntities {
		tasks = append(tasks, TaskEntityToModel(taskEntity))
	}
	return tasks, nil
}

func (pg *pgConnection) CreateTask(ctx context.Context, task models.CreateTask, tags []models.Tag) (models.Task, error) {
	checkListEntities := make([]entities.CheckList, 0)
	tagsEntities := make([]entities.Tag, 0)

	var checkList []models.CheckList
	for _, listItem := range task.CheckList {
		newCheckList := models.CheckList{
			Name:        listItem.Name,
			Description: listItem.Description,
		}
		checkList = append(checkList, newCheckList)
	}

	for _, listItem := range checkList {
		checkListEntities = append(checkListEntities, CheckListModelToEntity(listItem))
	}

	for _, tag := range tags {
		tagsEntities = append(tagsEntities, TagModelToEntity(tag))
	}

	taskEntity := entities.Task{
		Name:        task.Name,
		Description: task.Description,
		Deadline:    task.Deadline,
		CreatedAt:   time.Now().Unix(),
		Done:        false,
		CheckList:   checkListEntities,
		Tags:        tagsEntities,
	}

	tx := pg.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to create task")
		}
	}()
	if err := tx.Create(&taskEntity).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to create task: %v", err)
		return models.Task{}, pg.convertError(err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to create task: %v", err)
		return models.Task{}, pg.convertError(err)
	}

	return TaskEntityToModel(taskEntity), nil
}

func (pg *pgConnection) UpdateTask(ctx context.Context, id string, task models.UpdateTask) error {
	tx := pg.db.WithContext(ctx).Begin()
	query := tx.Model(&entities.Task{}).Where("id = ?", id)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to update task")
		}
	}()

	if task.Name != "" {
		if err := query.Update("name", task.Name).Error; err != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: failed to update task: %v", err)
			return pg.convertError(err)
		}
	}
	if task.Description != "" {
		if err := query.Update("description", task.Description).Error; err != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: failed to update task: %v", err)
			return pg.convertError(err)
		}
	}
	if task.Deadline != 0 {
		if err := query.Update("deadline", task.Deadline).Error; err != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: failed to update task: %v", err)
			return pg.convertError(err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to update task: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) AddCheckListItemToTask(ctx context.Context, task models.Task, newCheckList models.CreateCheckList) error {
	taskEntity := TaskModelToEntity(task)
	tx := pg.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to add checklist to task")
		}
	}()

	if err := tx.Model(&taskEntity).Association("CheckList").Append(&entities.CheckList{Name: newCheckList.Name, Description: newCheckList.Description}); err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to add checklist to task: %v", err)
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to add checklist to task: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) AddTagToTask(ctx context.Context, task models.Task, tag models.Tag) error {
	taskEntity := TaskModelToEntity(task)
	tagEntity := TagModelToEntity(tag)
	tx := pg.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to add tag to task")
		}
	}()

	if err := tx.Model(&taskEntity).Association("Tags").Append(&tagEntity); err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to add tag to task: %v", err)
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to add tag to task: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) CompleteTask(ctx context.Context, id string) error {
	tx := pg.db.WithContext(ctx).Begin()
	query := tx.Model(&entities.Task{}).Where("id = ?", id)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to complete task")
		}
	}()

	if err := query.Update("done", true).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to complete task: %v", err)
		return pg.convertError(err)
	}

	if err := query.Update("finished_at", time.Now().Unix()).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to complete task: %v", err)
		return pg.convertError(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to complete task: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) CompleteCheckListItem(ctx context.Context, id string) error {
	tx := pg.db.WithContext(ctx).Begin()
	query := tx.Model(&models.CheckList{}).Where("id = ?", id)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to complete check list")
		}
	}()

	if err := query.Update("done", true).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to complete check list: %v", err)
		return pg.convertError(err)
	}

	if err := query.Update("finished_at", time.Now().Unix()).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to complete check list: %v", err)
		return pg.convertError(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to complete check list: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) DeleteCheckListItemFromTask(ctx context.Context, task models.Task, checkList models.CheckList) error {
	taskEntity := TaskModelToEntity(task)
	checkListEntity := CheckListModelToEntity(checkList)

	tx := pg.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to delete checklist from task")
		}
	}()

	if err := tx.Model(&taskEntity).Association("CheckList").Delete(&checkListEntity); err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete checklist from task: %v", err)
		return err
	}
	// if err := tx.Model(&checkListEntity).Where("id = ?", c).Update("task_id", 0).Error; err != nil {
	// 	tx.Rollback()
	// 	pg.log.Errorw("postgres: failed to delete checklist from task: %v", err)
	// 	return err
	// }
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete checklist from task: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) DeleteTagFromTask(ctx context.Context, task models.Task, tag models.Tag) error {
	taskEntity := TaskModelToEntity(task)
	tagEntity := TagModelToEntity(tag)
	tx := pg.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to delete tag from task")
		}
	}()

	if err := tx.Model(&taskEntity).Association("Tags").Delete(&tagEntity); err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete tag from task: %v", err)
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete tag from task: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) DeleteTask(ctx context.Context, task models.Task) error {
	taskEntity := TaskModelToEntity(task)

	tx := pg.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to delete task")
		}
	}()
	if err := tx.Delete(&taskEntity).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete task: %v", err)
		return pg.convertError(err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete task: %v", err)
		return pg.convertError(err)
	}

	return nil
}
