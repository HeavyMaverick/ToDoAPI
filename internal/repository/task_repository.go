package repository

import (
	"ToDoApi/internal/model"
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("Task not found")
)

type TaskRepository interface {
	GetAll() ([]model.Task, error)
	GetById(id int) (*model.Task, error)
	Create(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
}

type InMemoryTaskRepository struct {
	tasks  []model.Task
	nextId int
	mu     sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make([]model.Task, 0),
		nextId: 0,
	}
}

func (r *InMemoryTaskRepository) GetAll() ([]model.Task, error) {
	return r.tasks, nil
}
func (r *InMemoryTaskRepository) GetById(id int) (*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			return &r.tasks[i], nil
		}
	}
	return nil, ErrNotFound
}

func (r *InMemoryTaskRepository) Create(task *model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nextId++
	task.ID = r.nextId
	task.CreatedAt = time.Now()
	r.tasks = append(r.tasks, *task)
	return nil
}
func (r *InMemoryTaskRepository) Update(id int, updatedTask *model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, task := range r.tasks {
		if task.ID == id {
			updatedTask.CreatedAt = task.CreatedAt
			updatedTask.ID = id
			r.tasks[i] = *updatedTask
			return nil
		}
	}
	return ErrNotFound
}

func (r *InMemoryTaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for index, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:index], r.tasks[index+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
