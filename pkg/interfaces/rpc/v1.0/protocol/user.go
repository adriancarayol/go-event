package protocol

import uuid "github.com/google/uuid"

type User struct {
	ID uuid.UUID
	Email string
}