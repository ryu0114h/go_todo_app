package service

import (
	"context"
	"fmt"

	"github.com/ryu0114h/go_todo_app/auth"
	"github.com/ryu0114h/go_todo_app/entity"
	"github.com/ryu0114h/go_todo_app/store"
)

type UpdateTask struct {
	DB   store.Execer
	Repo TaskUpdater
}

func (u *UpdateTask) UpdateTask(ctx context.Context, taskId int64, title, status string) (*entity.Task, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	isValidStatus := entity.ValidStatus(status)
	if !isValidStatus {
		status = string(entity.TaskStatusTodo)
	}
	t := &entity.Task{
		ID:     entity.TaskID(taskId),
		UserID: id,
		Title:  title,
		Status: entity.TaskStatus(status),
	}
	err := u.Repo.UpdateTask(ctx, u.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to update: %w", err)
	}
	return t, nil
}
