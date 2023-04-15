package service

import (
	"context"
	"fmt"

	"github.com/ryu0114h/go_todo_app/auth"
	"github.com/ryu0114h/go_todo_app/entity"
	"github.com/ryu0114h/go_todo_app/store"
)

type DeleteTask struct {
	DB   store.Execer
	Repo TaskDeleter
}

func (d *DeleteTask) DeleteTask(ctx context.Context, taskID entity.TaskID) error {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("user_id not found")
	}
	err := d.Repo.DeleteTask(ctx, d.DB, id, taskID)
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}
	return nil
}
