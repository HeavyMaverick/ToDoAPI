package service

import (
	"ToDoApi/internal/model"
	"ToDoApi/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserWithId(id int) (*model.User, error)
	GetUserWithUsername(username string) *model.User
	CreateUser(username, pass, email string) (*model.User, error)
	DeleteUserWithId(id int) error
	DeleteUserWithUsername(username string) error
	UpdateUser(id int, username, pass string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetUserWithId(id int) (*model.User, error) {
	return s.userRepo.GetById(id)
}

func (s *userService) GetUserWithUsername(username string) (*model.User, error) {
	return s.userRepo.GetByUsername(username)
}

func (s *userService) CreateUser(username, pass, email string) (*model.User, error) {
	return s.userRepo.Create(username, pass, email)
}

func (s *userService) DeleteUserWithId(id int) error {
	return s.userRepo.DeleteUserById(id)
}

func (s *userService) DeleteUserWithUsername(username string) error {
	return s.userRepo.DeleteUserByUsername(username)
}

func (s *userService) UpdateUser(id int, username, pass string) (*model.User, error)
