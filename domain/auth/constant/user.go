package constant

type UserRole string

func (ur UserRole) String() string {
	return string(ur)
}

const (
	RoleAdmin    UserRole = "admin"
	RoleStreamer UserRole = "streamer"
)

type UserStatus string

func (us UserStatus) String() string {
	return string(us)
}

const (
	StatusActive UserStatus = "active"
	StatusBanned UserStatus = "banned"
)
