package repository

import (
	"ToDoApi/internal/model"

	"gorm.io/gorm"
)

type PostgresTaskRepository struct {
	db *gorm.DB
}

func NewPostgresTaskRepository(db *gorm.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}
func (r *PostgresTaskRepository) GetById(id int) (*model.Task, error)
func (r *PostgresTaskRepository) Create(task *model.Task) error
func (r *PostgresTaskRepository) Update(id int, task *model.Task) error
func (r *PostgresTaskRepository) Delete(id int) error
