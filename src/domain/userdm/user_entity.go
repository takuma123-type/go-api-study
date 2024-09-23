package userdm

import (
	"fmt"
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type User struct {
	ID        UserID           `gorm:"column:id"`
	FirstName string           `gorm:"column:first_name"`
	LastName  string           `gorm:"column:last_name"`
	CreatedAt shared.CreatedAt `gorm:"column:created_at"`
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetCreatedAt() shared.CreatedAt {
	return u.CreatedAt
}

var (
	firstNameLength = 30
	lastNameLength  = 30
)

func newUser(id UserID, first, last string, createdAt shared.CreatedAt) (*User, error) {
	if first == "" {
		return nil, smperr.BadRequest("first name must not be empty")
	}
	if last == "" {
		return nil, smperr.BadRequest("last name must not be empty")
	}

	if l := utf8.RuneCountInString(first); l > firstNameLength {
		return nil, smperr.BadRequest(
			fmt.Sprintf("first name must be less than %d characters", firstNameLength),
		)
	}
	if l := utf8.RuneCountInString(last); l > lastNameLength {
		return nil, smperr.BadRequest(
			fmt.Sprintf("last name must be less than %d characters", lastNameLength),
		)
	}

	return &User{
		ID:        id,
		FirstName: first,
		LastName:  last,
		CreatedAt: createdAt,
	}, nil
}

func (u *User) UpdateUser(first, last string) error {
	if first == "" {
		return smperr.BadRequest("first name must not be empty")
	}
	if last == "" {
		return smperr.BadRequest("last name must not be empty")
	}

	if l := utf8.RuneCountInString(first); l > firstNameLength {
		return smperr.BadRequest(
			fmt.Sprintf("first name must be less than %d characters", firstNameLength),
		)
	}
	if l := utf8.RuneCountInString(last); l > lastNameLength {
		return smperr.BadRequest(
			fmt.Sprintf("last name must be less than %d characters", lastNameLength),
		)
	}

	u.FirstName = first
	u.LastName = last

	return nil
}
