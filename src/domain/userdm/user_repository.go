package userdm

import (
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, userID UserID) (*User, error)
}
