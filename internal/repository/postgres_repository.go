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
	var existingTask model.Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}
	updates := map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"completed":   task.Completed,
		"user_id":     task.UserID,
	}
	result := r.db.Model(&model.Task{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if err := r.db.First(task, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresTaskRepository) Delete(id int) error {
	result := r.db.Where("id=?", id).Delete(&model.Task{})
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return result.Error
}
