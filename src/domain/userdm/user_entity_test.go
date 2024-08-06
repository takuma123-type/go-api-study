package userdm

import (
	"testing"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

func TestNewUser(t *testing.T) {
	userID := NewUserID()
	createdAt := shared.NewCreatedAt()

	user, err := newUser(userID, "John", "Doe", createdAt)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if user.ID() != userID {
		t.Errorf("Expected userID %v, got %v", userID, user.ID())
	}

	if user.FirstName() != "John" {
		t.Errorf("Expected firstName John, got %v", user.FirstName())
	}

	if user.LastName() != "Doe" {
		t.Errorf("Expected lastName Doe, got %v", user.LastName())
	}
}
