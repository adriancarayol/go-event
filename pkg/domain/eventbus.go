package domain

import (
	"github.com/adriancarayol/go-event/pkg/dao"
)

type EventBus interface {
	Publish(event dao.Event) error
}
