package model

import (
	"encoding/json"
	"github.com/adriancarayol/go-event/pkg/domain"
	"github.com/adriancarayol/go-event/pkg/domain/commands"
	"github.com/adriancarayol/go-event/pkg/domain/events"
	"github.com/google/uuid"
	"time"
)

type User struct {
	domain.BaseAggregate
	Username     string
	Email        string
	PasswordHash string
}

func (user *User) GetID() uuid.UUID {
	return user.ID
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) GetUsername() string {
	return user.Username
}

func (user *User) IncrementVersion() {
	user.Version += 1
}

func (user *User) UpdateUpdatedAt(t time.Time) {
	user.UpdatedAt = t
}

func (user *User) AggregateType() string {
	return "User"
}

func (user *User) ApplyChange(event events.Event) error {
	switch event.Type {
	case events.Created:
		payload := events.UserCreatedData{}
		jsonString, err := json.Marshal(event.Data)

		if err != nil {
			return err
		}

		err = json.Unmarshal(jsonString, &payload)

		if err != nil {
			return err
		}

		user.Username = payload.Username
		user.Email = payload.Email
		user.PasswordHash = payload.Password
		user.ID = event.AggregateID
	}

	return nil
}

func (user *User) HandleCommand(command commands.Command) error {
	eventID, err := uuid.NewRandom()

	if err != nil {
		return nil
	}

	event := events.Event{
		ID:            eventID,
		AggregateID:   user.ID,
		AggregateType: user.AggregateType(),
	}

	switch c := command.(type) {
	case commands.CreateUser:
		event.AggregateID = c.AggregateID
		event.Type = events.Created
		data := make(map[string]interface{})
		data["username"] = c.Username
		data["email"] = c.Email
		data["password"] = c.Password
		event.Data = data
	}

	err = user.ApplyChange(event)

	if err != nil {
		return err
	}

	user.Events = append(user.Events, event)

	return nil
}
