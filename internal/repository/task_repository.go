package repository

import (
	"ToDoApi/internal/model"
	"sync"
	"time"
)

type TaskRepository interface {
	GetAll() ([]model.Task, error)
	GetById(id int) (*model.Task, error)
	Create(model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
}

type InMemoryTaskRepository struct {
	tasks  []model.Task
	nextId int
	mu     sync.RWMutex
}

func (r *InMemoryTaskRepository) Create(task model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nextId++
	task.ID = r.nextId
	task.CreatedAt = time.Now()
	r.tasks = append(r.tasks, task)
	return nil
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make([]model.Task, 0),
		nextId: 0,
	}
}
