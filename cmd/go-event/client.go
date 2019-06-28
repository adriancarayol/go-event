package main

import (
	"context"
	"encoding/json"
	"github.com/adriancarayol/go-event/config/mongo"
	"github.com/adriancarayol/go-event/config/rabbit"
	"github.com/adriancarayol/go-event/pkg/dao"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Data struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	rabbit.Init()
	mongo.Init()

	amqpConnection := rabbit.Get()
	ch, err := amqpConnection.Channel()

	if err != nil {
		log.Fatalln("Fail creating channel")
	}

	channel, err := ch.Consume("user_events", "", false, false, false, false, nil)

	forever := make(chan bool)

	collection := mongo.Get().Database("testing").Collection("numbers")

	go func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		for d := range channel {
			log.Printf("Received a message: %s", d.Body)
			var event dao.Event
			var data Data

			err = json.Unmarshal(d.Body, &event)

			if err != nil {
				log.Printf("Error: %s", err)
			}

			err = json.Unmarshal([]byte(event.Data), &data)

			if err != nil {
				log.Printf("Error: %s", err)
			}

			res, err := collection.InsertOne(ctx, bson.M{"id": event.AggregateID.String(), "password": data.Password, "username": data.Username, "email": data.Email})

			if err != nil {
				log.Printf("Error: %s", err)
			}

			log.Printf("Inserted into mongo: %s", res)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
