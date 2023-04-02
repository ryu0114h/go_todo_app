package store

import (
	"errors"

	"github.com/ryu0114h/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{
		Tasks: map[entity.TaskID]*entity.Task{},
	}
	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) Get(id entity.TaskID) (*entity.Task, error) {
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
