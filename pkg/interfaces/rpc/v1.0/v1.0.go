package v1_0

import (
	"github.com/adriancarayol/go-event/pkg/interfaces/registry"
	"github.com/adriancarayol/go-event/pkg/interfaces/rpc/v1.0/protocol"
	"github.com/adriancarayol/go-event/pkg/usecases"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecases.UserUseCase)))
}
