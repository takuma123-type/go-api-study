package userusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/useroutput"
)

type GetUserListUsecase struct {
	userRepository userdm.UserRepository
}

func NewGetUserList(userRepo userdm.UserRepository) *GetUserListUsecase {
	return &GetUserListUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserListUsecase) Exec(ctx context.Context) (*useroutput.UserList, error) {
	users, err := use.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	userItems := make([]*useroutput.UserItem, len(users))
	for i, user := range users {
		userItems[i] = &useroutput.UserItem{
			ID:        user.ID.String(),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
		}
	}
	return &useroutput.UserList{Users: userItems}, nil
}
