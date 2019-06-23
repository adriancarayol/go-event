package domain

import (
	"github.com/adriancarayol/go-event/pkg/domain/events"
	"github.com/google/uuid"
	"time"
)

type Aggregate interface {
	GetID() uuid.UUID
	IncrementVersion()
	UpdateUpdatedAt(time.Time)
	AggregateType() string
}

type BaseAggregate struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at"`
	Version   uint64         `json:"version"`
	Events    []events.Event `json:"events"`
}

func (agg *BaseAggregate) GetID() uuid.UUID {
	return agg.ID
}

func (agg *BaseAggregate) IncrementVersion() {
	agg.Version += 1
}

func (agg *BaseAggregate) UpdateUpdatedAt(t time.Time) {
	agg.UpdatedAt = t
}
