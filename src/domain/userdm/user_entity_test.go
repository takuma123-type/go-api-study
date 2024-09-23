package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		lastName  string
		expectErr bool
	}{
		{
			name:      "Valid user",
			firstName: "John",
			lastName:  "Doe",
			expectErr: false,
		},
		{
			name:      "Empty first name",
			firstName: "",
			lastName:  "Doe",
			expectErr: true,
		},
		{
			name:      "Empty last name",
			firstName: "John",
			lastName:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID := NewUserID()
			createdAt := shared.NewCreatedAt()

			user, err := newUser(userID, tt.firstName, tt.lastName, createdAt)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, userID, user.ID)
				assert.Equal(t, tt.firstName, user.FirstName)
				assert.Equal(t, tt.lastName, user.LastName)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name         string
		firstName    string
		lastName     string
		newFirstName string
		newLastName  string
		expectErr    bool
	}{
		{
			name:         "Valid user update",
			firstName:    "John",
			lastName:     "Doe",
			newFirstName: "Jane",
			newLastName:  "Smith",
			expectErr:    false,
		},
		{
			name:         "Empty new first name",
			firstName:    "John",
			lastName:     "Doe",
			newFirstName: "",
			newLastName:  "Smith",
			expectErr:    true,
		},
		{
			name:         "Empty new last name",
			firstName:    "John",
			lastName:     "Doe",
			newFirstName: "Jane",
			newLastName:  "",
			expectErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID := NewUserID()
			createdAt := shared.NewCreatedAt()

			user, err := newUser(userID, tt.firstName, tt.lastName, createdAt)
			assert.NoError(t, err)

			err = user.UpdateUser(tt.newFirstName, tt.newLastName)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newFirstName, user.FirstName)
				assert.Equal(t, tt.newLastName, user.LastName)
			}
		})
	}
}
