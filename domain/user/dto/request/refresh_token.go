package request

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
