package userusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/userinput"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/useroutput"
)

type CreateUserUsecase struct {
	userRepository userdm.UserRepository
}

func NewCreateUser(userRepo userdm.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepository: userRepo,
	}
}

func (use *CreateUserUsecase) Exec(ctx context.Context, in *userinput.CreateUserInput) (*useroutput.CreateUserOutput, error) {

	user, err := userdm.GenWhenCreate(in.FirstName, in.LastName)
	if err != nil {
		return nil, err
	}

	if err := use.userRepository.Store(ctx, user); err != nil {
		return nil, err
	}

	return &useroutput.CreateUserOutput{
		ID:        user.ID().String(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		CreatedAt: user.CreatedAt().Value(),
	}, nil
}
