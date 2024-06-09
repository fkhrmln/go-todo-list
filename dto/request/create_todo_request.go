package request

type CreateTodoRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body"`
	Deadline int    `json:"deadline" validate:"required"`
}
