package models

// User represents a user in the system
// swagger:model User
type User struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}
