package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserID(t *testing.T) {
	tests := []struct {
		name      string
		val       string
		expectErr bool
		expected  UserID
	}{
		{
			name:      "Valid non-empty UserID",
			val:       "123456",
			expectErr: false,
			expected:  "123456",
		},
		{
			name:      "Empty UserID",
			val:       "",
			expectErr: true,
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID, err := NewUserIDByVal(tt.val)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, userID)
			}
		})
	}

	t.Run("NewUserID should create a valid UserID", func(t *testing.T) {
		userID := NewUserID()
		assert.NotEmpty(t, userID.String(), "Expected a non-empty UserID")
	})

	t.Run("Equal should return true for identical UserIDs", func(t *testing.T) {
		val := "123456"
		userID1, _ := NewUserIDByVal(val)
		userID2, _ := NewUserIDByVal(val)
		assert.True(t, userID1.Equal(userID2), "Expected Equal to return true for identical UserIDs")
	})

	t.Run("Equal should return false for different UserIDs", func(t *testing.T) {
		userID1 := NewUserID()
		userID2 := NewUserID()
		assert.False(t, userID1.Equal(userID2), "Expected Equal to return false for different UserIDs")
	})
}
