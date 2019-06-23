package service

import (
	"fmt"
	"github.com/adriancarayol/go-event/pkg/domain/repository"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	repo repository.UserRepository
	eventStore repository.EventStore
}

func NewUserService(repo repository.UserRepository, eventStore repository.EventStore) *UserService {
	return &UserService{
		eventStore: eventStore,
		repo: repo,
	}
}

func (s *UserService) CheckIfExistEmail(email string) error {
	user, err := s.repo.FindByEmail(email)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	if user != nil {
		return fmt.Errorf("user with email: %s already exists", email)
	}

	return nil
}

func (s *UserService) CheckIfExistUsername(username string) error {
	user, err := s.repo.FindByUsername(username)

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	if user != nil {
		return fmt.Errorf("user with username: %s already exists", username)
	}

	return nil
}