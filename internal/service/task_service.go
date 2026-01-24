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

type taskService struct {
	rep repository.TaskRepository
}

func (s *taskService) CreateTask(task *model.Task) error {
	if task.Title == "" {
		return errors.New("Error: empty title. Can't be empty")
	}
	return s.rep.Create(*task)
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

func (s *taskService) UpdateTask(id int, model *model.Task) error {
	return s.rep.Update(id, model)
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{rep: repo}
}
