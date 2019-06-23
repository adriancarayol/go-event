package db

import (
	"github.com/adriancarayol/go-event/pkg/domain/events"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("postgres", "host=localhost port=26257 user=goevent dbname=goevent sslmode=disable")

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	db.AutoMigrate(&events.Event{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
