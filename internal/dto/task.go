package dto

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=60"`
	Description string `json:"description" binding:"max=1000"`
	UserID      int    `json:"user_id" binding:"required,min=1"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=60"`
	Description string `json:"description" binding:"max=1000"`
	Completed   bool   `json:"completed"`
}
