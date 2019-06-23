# go-event
DDD + CQRS + Event sourcing example with GoLang + PostgreSQL (Cockroach).

* Install postgresql and cockroach
* Install go and dep.

## Steps
* dep ensure inside go-event root folder.
* Start psql service
* Execute `cockroach start --insecure --listen-addr=localhost`
* Execute go run cmd/go-event/main.go
