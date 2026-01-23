package request

type UserLogoutPayload struct {
	AccessToken string `json:"access_token" validate:"required"`
}
