package userdm

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

type User struct {
	id        UserID
	firstName string
	lastName  string
	createdAt shared.CreatedAt
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) CreatedAt() shared.CreatedAt {
	return u.createdAt
}

var (
	firstNameLength = 30
	lastNameLength  = 30
)

func newUser(id UserID, first, last string, createdAt shared.CreatedAt) (*User, error) {
	if first == "" {
		return nil, errors.New("first name must not be empty")
	}
	if last == "" {
		return nil, errors.New("last name must not be empty")
	}

	if l := utf8.RuneCountInString(first); l > firstNameLength {
		return nil, fmt.Errorf("first name must be less than %d characters", firstNameLength)
	}
	if l := utf8.RuneCountInString(last); l > lastNameLength {
		return nil, fmt.Errorf("last name must be less than %d characters", lastNameLength)
	}

	return &User{
		id:        id,
		firstName: first,
		lastName:  last,
		createdAt: createdAt,
	}, nil
}
