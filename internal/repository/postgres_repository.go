package repository

import (
	"ToDoApi/internal/model"
	"errors"

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
	result := r.db.Order("id").Find(&tasks)
	return tasks, result.Error
}

func (r *PostgresTaskRepository) GetById(id int) (*model.Task, error) {
	var task model.Task
	result := r.db.First(&task, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &task, result.Error
}

func (r *PostgresTaskRepository) Create(task *model.Task) error {
	result := r.db.Create(task)
	return result.Error
}

func (r *PostgresTaskRepository) Update(id int, task *model.Task) error {
	result := r.db.Model(&model.Task{}).Where("id=?", id).Updates(map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"completed":   task.Completed,
	})
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	if err := r.db.First(task, id).Error; err != nil {
		return err
	}
	return result.Error
}

func (r *PostgresTaskRepository) Delete(id int) error {
	result := r.db.Where("id=?", id).Delete(&model.Task{})
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return result.Error
}
