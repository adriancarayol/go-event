package usecases

import (
	"github.com/adriancarayol/go-event/pkg/dao"
	"github.com/adriancarayol/go-event/pkg/domain"
	"github.com/adriancarayol/go-event/pkg/domain/commands"
	"github.com/adriancarayol/go-event/pkg/domain/model"
	"github.com/adriancarayol/go-event/pkg/domain/repository"
	"github.com/adriancarayol/go-event/pkg/domain/service"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(email, username, password string) error
}

type userUseCase struct {
	repo       repository.UserRepository
	eventStore repository.EventStore
	eventBus   domain.EventBus
	service    *service.UserService
}

func NewUserUsecase(repo repository.UserRepository, eventStore repository.EventStore, eventBus domain.EventBus, service *service.UserService) *userUseCase {
	return &userUseCase{
		repo:       repo,
		eventStore: eventStore,
		eventBus:   eventBus,
		service:    service,
	}
}

func (u *userUseCase) RegisterUser(email, username, password string) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return err
	}

	/*if err := u.service.CheckIfExistEmail(email); err != nil {
		return err
	}

	if err := u.service.CheckIfExistUsername(username); err != nil {
		return err
	}*/

	user := &model.User{}

	userID, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	createUserCommand := commands.CreateUser{
		Username: username,
		Email:    email,
		Password: string(passwordHash),
	}

	createUserCommand.AggregateID = userID
	createUserCommand.AggregateType = user.AggregateType()

	err = user.HandleCommand(createUserCommand)

	if err != nil {
		return err
	}

	err = u.eventStore.Store(user.Events)

	if err != nil {
		return err
	}

	for _, event := range user.Events {
		var eventDao dao.Event

		payload, err := event.Data.Value()

		if err != nil {
			return err
		}

		data := payload.([]byte)
		eventDao.ID = event.ID
		eventDao.AggregateID = event.AggregateID
		eventDao.AggregateType = event.AggregateType
		eventDao.Version = event.Version
		eventDao.Type = string(event.Type)
		eventDao.Data = string(data)

		if err := u.eventBus.Publish(eventDao); err != nil {
			return err
		}
	}

	return nil
}
