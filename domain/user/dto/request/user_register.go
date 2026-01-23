package request

type UserRegisterPayload struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8"`
	Fullname string `json:"fullname" validate:"required,min=3,max=100"`
}
