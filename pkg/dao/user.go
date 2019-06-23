package dao

import (
	"github.com/adriancarayol/go-event/pkg/domain/model"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Email string
	Username string
}

func toUser(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID: user.GetID(),
			Email: user.GetEmail(),
			Username: user.GetUsername(),
		}
	}
	return res
}