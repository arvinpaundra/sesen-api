package dto

type SetupOverlay struct {
	UserId string `json:"user_id" validate:"required,uuid"`
}
