package repository

import (
	"ToDoApi/internal/model"
	"errors"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetById(id int) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Exists(id int) (bool, error)
	Create(username, pass, email string) (*model.User, error)
	DeleteUserByUsername(username string) error
	DeleteUserById(id int) error
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	result := r.db.Order("id").Find(&users)
	return users, result.Error
}

func (r *PostgresUserRepository) GetById(id int) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, username)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *PostgresUserRepository) Exists(id int) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresUserRepository) Create(username, pass, email string) (*model.User, error) {
	var user model.User
	return &user, nil
}
