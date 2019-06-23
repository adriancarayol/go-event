package repository

import (
	"github.com/adriancarayol/go-event/pkg/domain/events"
	"github.com/google/uuid"
)

type EventStore interface {
	Store([]events.Event) error
	Load(uuid uuid.UUID) ([]events.Event, error)
}
