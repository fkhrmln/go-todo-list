package request

type UpdateTodoRequest struct {
	ID          string `json:"id" validate:"required"`
	UserID      string `json:"user_id" validate:"required"`
	Title       string `json:"title" gorm:"omitempty"`
	Body        string `json:"body" gorm:"omitempty"`
	Deadline    int    `json:"deadline" gorm:"omitempty"`
	IsCompleted bool   `json:"is_completed"`
}
