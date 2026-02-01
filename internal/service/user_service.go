package service

import (
	"ToDoApi/internal/model"
	"ToDoApi/internal/repository"

	"golang.org/x/crypto/bcrypt"
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

// func NewUserService(userRepo repository.UserRepository) UserService {
// 	return &userService{userRepo: userRepo}
// }

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

func (s *userService) UpdateUser(id int, username, pass string) (*model.User, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	if username != "" {
		user.Username = username
	}
	if pass != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = string(hashedPassword)
	}
	return user, nil
}
