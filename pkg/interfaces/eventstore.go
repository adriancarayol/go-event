package interfaces

import (
	"fmt"
	"github.com/adriancarayol/go-event/config/db"
	"github.com/adriancarayol/go-event/pkg/domain/events"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type eventStore struct {
	db *gorm.DB
}

type JSONB map[string]interface{}


func NewEventStore() *eventStore {
	return &eventStore{
		db: db.GetDB(),
	}
}

func (e *eventStore) Store(es []events.Event) error {
	tx := e.db.Begin()

	for _, e := range es {
		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("error while saving event: %s", err)
		}
	}

	return tx.Commit().Error

}

func (e *eventStore) Load(uuid uuid.UUID) ([]events.Event, error) {
	var es []events.Event

	return es, nil
}
