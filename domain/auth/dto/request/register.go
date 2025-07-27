package request

type UserRegister struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=8"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Fullname string `json:"fullname" validate:"required,min=3,max=100"`
}
