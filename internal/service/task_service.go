package service

import (
	"ToDoApi/internal/model"
	"ToDoApi/internal/repository"
	"errors"
)

type TaskService interface {
	GetAllTasks() ([]model.Task, error)
	GetTask(id int) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(id int, task *model.Task) error
	DeleteTask(id int) error
}

var (
	ErrEmptyTitle   = errors.New("Title cannot be empty")
	ErrTitleTooLong = errors.New("Title too long (max 100 chars)")
)

type taskService struct {
	rep repository.TaskRepository
}

func (s *taskService) CreateTask(task *model.Task) error {
	if task.Title == "" {
		return ErrEmptyTitle
	}
	if len(task.Title) > 100 {
		return ErrTitleTooLong
	}
	return s.rep.Create(task)
}

func (s *taskService) GetAllTasks() ([]model.Task, error) {
	return s.rep.GetAll()
}

func (s *taskService) DeleteTask(id int) error {
	return s.rep.Delete(id)
}

func (s *taskService) GetTask(id int) (*model.Task, error) {
	return s.rep.GetById(id)
}

func (s *taskService) UpdateTask(id int, task *model.Task) error {
	if task.Title == "" {
		return ErrEmptyTitle
	}
	if len(task.Title) > 100 {
		return ErrTitleTooLong
	}
	return s.rep.Update(id, task)
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{rep: repo}
}
