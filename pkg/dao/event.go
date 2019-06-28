package dao

import (
	"github.com/google/uuid"
)

type Event struct {
	ID            uuid.UUID
	AggregateID   uuid.UUID
	AggregateType string
	Type          string
	Version       uint64
	Data          string
}
