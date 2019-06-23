package repository

import "github.com/adriancarayol/go-event/pkg/domain/model"

type UserRepository interface {
	Save(*model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}