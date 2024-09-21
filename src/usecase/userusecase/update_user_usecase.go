package userusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/userinput"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/useroutput"
)

type UpdateUserUsecase struct {
	userRepository userdm.UserRepository
}

func NewUpdateUser(userRepo userdm.UserRepository) *UpdateUserUsecase {
	return &UpdateUserUsecase{
		userRepository: userRepo,
	}
}

func (use *UpdateUserUsecase) Exec(ctx context.Context, in *userinput.UpdateUserInput) (*useroutput.UpdateUserOutput, error) {
	user, err := use.userRepository.FindByID(ctx, userdm.UserID(in.ID))
	if err != nil {
		return nil, err
	}

	if err := user.UpdateUser(in.FirstName, in.LastName); err != nil {
		return nil, err
	}

	if err := use.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return &useroutput.UpdateUserOutput{
		ID:        user.ID.String(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		UpdatedAt: user.GetCreatedAt().Value(),
	}, nil
}
