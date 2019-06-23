package v1_0

import (
	"context"
	"github.com/adriancarayol/go-event/pkg/interfaces/rpc/v1.0/protocol"
	"github.com/adriancarayol/go-event/pkg/usecases"
)

type userService struct {
	userUseCase usecases.UserUseCase
}

func NewUserService(userUseCase usecases.UserUseCase) *userService {
	return &userService{
		userUseCase: userUseCase,
	}
}

func (u *userService) RegisterUser(ctx context.Context, in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error) {
	if err := u.userUseCase.RegisterUser(in.GetEmail()); err != nil {
		return &protocol.RegisterUserResponseType{}, err
	}
	return &protocol.RegisterUserResponseType{}, nil
}