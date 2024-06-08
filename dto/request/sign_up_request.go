package request

type AuthRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required,max=100"`
}
