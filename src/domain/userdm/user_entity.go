package userdm

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

type User struct {
	ID        UserID           `gorm:"column:id"`
	FirstName string           `gorm:"column:first_name"`
	LastName  string           `gorm:"column:last_name"`
	CreatedAt shared.CreatedAt `gorm:"column:created_at"`
	UpdatedAt time.Time        `gorm:"column:updated_at"`
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

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) Update(firstName, lastName string) {
	u.FirstName = firstName
	u.LastName = lastName
	u.UpdatedAt = time.Now()
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
		ID:        id,
		FirstName: first,
		LastName:  last,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
	}, nil
}
