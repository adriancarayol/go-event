# go-event
DDD + CQRS + Event sourcing example with GoLang + PostgreSQL as Store Event (Cockroach) and MongoDB as Read Model.

* Install postgresql and cockroach
* Install go and dep


## TODO
* Implement reconnect funtionality to RabbitMQ
* Create repository for MongoDB
* Implement read model using MongoDB
* Implement Saga orchestrator for user registration

## Steps
* dep ensure inside go-event root folder
* Start psql service
* Execute `cockroach start --insecure --listen-addr=localhost`
* Execute go run cmd/go-event/main.go
* Start mongoDB service (see config/mongo)
* Start RabbitMQ service (see config/rabbitmq)
