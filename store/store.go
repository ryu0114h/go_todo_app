package store

import (
	"errors"

	"github.com/ryu0114h/go_todo_app/entity"
)

var (
	Tasks = TaskStore{
		Tasks: map[entity.TaskId]*entity.Task{},
	}
	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskId
	Tasks  map[entity.TaskId]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskId, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[ts.LastID] = t
	return ts.LastID, nil
}

func (ts *TaskStore) Get(id entity.TaskId) (*entity.Task, error) {
	if t, ok := ts.Tasks[id]; ok {
		return t, nil
	}
	return nil, ErrNotFound
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make(entity.Tasks, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
		// tasks = append(tasks, t)
	}
	return tasks
}
