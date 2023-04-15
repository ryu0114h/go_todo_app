package store

import (
	"context"
	"fmt"
	"time"

	"github.com/ryu0114h/go_todo_app/entity"
)

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
			(user_id, title, status, created, modified)
			VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.UserID, t.Title, t.Status,
		t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) UpdateTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Modified = time.Now()
	sql := `UPDATE task 
			SET title = ?,
				status = ?,
				modified = ?
			WHERE user_id = ?
				AND id = ?`
	result, err := db.ExecContext(ctx, sql, t.Title, t.Status, t.Modified, t.UserID, t.ID)
	if err != nil {
		return err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if id == 0 {
		return fmt.Errorf("no changed")
	}
	return nil
}

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, id entity.UserID,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
			id, user_id, title,
			status, created, modified
			FROM task
			WHERE user_id = ?;`
	if err := db.SelectContext(ctx, &tasks, sql, id); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) DeleteTask(
	ctx context.Context, db Execer, userID entity.UserID, taskId entity.TaskID,
) error {
	sql := `DELETE FROM task
			WHERE user_id = ?
				AND id = ?;
			`
	result, err := db.ExecContext(
		ctx, sql, userID, taskId,
	)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("no changed")
	}
	return nil
}
