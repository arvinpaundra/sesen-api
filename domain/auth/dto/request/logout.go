package request

type UserLogout struct {
	AccessToken string `json:"access_token" validate:"required"`
}
