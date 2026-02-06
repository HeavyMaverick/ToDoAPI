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
	rep      repository.TaskRepository
	userRepo repository.UserRepository
}

var (
	ErrEmptyTitle         = errors.New("title cannot be empty")
	ErrTitleTooLong       = errors.New("title too long (max 60 chars)")
	ErrTitleTooShort      = errors.New("title too short (min 1 char)")
	ErrDescriptionTooLong = errors.New("description too long")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrUserNotFound       = errors.New("user not found")
)

const (
	MaxTitleLength       = 60
	MaxDescriptionLength = 1000
	MinTitleLength       = 1
)

func (s *taskService) CreateTask(task *model.Task) error {
	err := s.validateTask(task)
	if err != nil {
		return err
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
	err := s.validateTask(task)
	if err != nil {
		return err
	}
	return s.rep.Update(id, task)
}

func NewTaskService(taskRepo repository.TaskRepository, userRepo repository.UserRepository) TaskService {
	return &taskService{
		rep:      taskRepo,
		userRepo: userRepo,
	}
}

func (s *taskService) validateTask(task *model.Task) error {
	if task.Title == "" {
		return ErrEmptyTitle
	}
	if len(task.Title) > MaxTitleLength {
		return ErrTitleTooLong
	}
	if len(task.Title) < MinTitleLength {
		return ErrTitleTooShort
	}
	if len(task.Description) > MaxDescriptionLength {
		return ErrDescriptionTooLong
	}
	if task.UserID <= 0 {
		return ErrInvalidUserID
	}
	if s.userRepo != nil {
		exists, err := s.userRepo.Exists(task.UserID)
		if err != nil {
			return err
		}
		if !exists {
			return ErrUserNotFound
		}
	}
	return nil
}
