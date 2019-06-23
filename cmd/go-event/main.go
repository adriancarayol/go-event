package main

import (
	"github.com/adriancarayol/go-event/config/db"
	"github.com/adriancarayol/go-event/pkg/interfaces/registry"
	"github.com/adriancarayol/go-event/pkg/usecases"
	"log"
)

func main() {
	db.Init()

	container, err := registry.NewContainer()
	if err != nil {
		log.Fatal("Cannot build the container")
	}

	userCase := container.Resolve("user-usecase").(usecases.UserUseCase)
	err = userCase.RegisterUser("foo@foo.com", "foousername", "foopassword")

	if err != nil {
		log.Fatalf("Error registering user: %s", err)
	}
}
