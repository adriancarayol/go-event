package registry

import (
	"github.com/adriancarayol/go-event/pkg/domain/service"
	"github.com/adriancarayol/go-event/pkg/interfaces"
	"github.com/adriancarayol/go-event/pkg/usecases"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	if err := builder.Add([]di.Def{
		{
			Name:  "user-usecase",
			Build: buildUserUseCase,
		},
	}...); err != nil {
		return nil, err
	}
	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUseCase(ctn di.Container) (interface{}, error) {
	repo := interfaces.NewUserRepository()
	eventStore := interfaces.NewEventStore()
	userService := service.NewUserService(repo, eventStore)
	return usecases.NewUserUsecase(repo, eventStore, userService), nil
}
