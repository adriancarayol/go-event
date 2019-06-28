package dao

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Email string
	Username string
}