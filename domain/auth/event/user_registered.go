package event

import (
	"encoding/json"

	"github.com/arvinpaundra/sesen-api/core/event"
)

const (
	UserRegisteredEventType = "user.registered"
)

// UserRegisteredEvent represents the event when a user registers
type UserRegisteredEvent struct {
	*event.BaseDomainEvent
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

func NewUserRegisteredEvent(userId, email, username, fullname string) *UserRegisteredEvent {
	return &UserRegisteredEvent{
		BaseDomainEvent: event.NewBaseDomainEvent(UserRegisteredEventType),
		UserId:          userId,
		Email:           email,
		Username:        username,
		Fullname:        fullname,
	}
}

// ToBytes serializes the event to JSON bytes
func (e *UserRegisteredEvent) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}

// GetUserID returns the user ID from the event payload
func (e *UserRegisteredEvent) GetUserID() string {
	return e.UserId
}

// GetEmail returns the email from the event payload
func (e *UserRegisteredEvent) GetEmail() string {
	return e.Email
}

// GetUsername returns the username from the event payload
func (e *UserRegisteredEvent) GetUsername() string {
	return e.Username
}

// GetFullname returns the fullname from the event payload
func (e *UserRegisteredEvent) GetFullname() string {
	return e.Fullname
}
