package events

const (
	Created = EventType("user:created")
)

type UserCreatedData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
