package events

import (
	"github.com/google/uuid"
)

type EventType string

type Event struct {
	ID            uuid.UUID   `json:"id" gorm:"type:uuid;primary_key;"`
	AggregateID   uuid.UUID   `json:"aggregate_id" gorm:"type:uuid"`
	AggregateType string      `json:"aggregate_type"`
	Type          EventType   `json:"type"`
	Version       uint64      `json:"version"`
	Data          PropertyMap `json:"-" gorm:"type:jsonb;column:data"`
}
