package commands

import "github.com/google/uuid"

type Command interface {
	GetAggregateID() uuid.UUID
	GetAggregateType() string
	GetCommandType() string
}

type BaseCommand struct {
	Type          string
	AggregateID   uuid.UUID
	AggregateType string
}

func (c BaseCommand) GetAggregateID() uuid.UUID {
	return c.AggregateID
}

func (c BaseCommand) GetAggregateType() string {
	return c.AggregateType
}

func (c BaseCommand) GetCommandType() string {
	return c.Type
}
