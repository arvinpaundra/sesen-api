package request

type CheckSessionPayload struct {
	AccessToken string `json:"access_token" binding:"required"`
}
