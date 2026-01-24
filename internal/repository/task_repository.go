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
	Create(model.Task) error
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
	for _, task := range r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return &model.Task{}, ErrNotFound
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
func (r *InMemoryTaskRepository) Update(id int, diffTask model.Task) error {
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
