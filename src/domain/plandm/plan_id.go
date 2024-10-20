package plandm

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// PlanID is a wrapper around UUID for plan identifiers
type PlanID uuid.UUID

// NewPlanID generates a new PlanID
func NewPlanID() PlanID {
	return PlanID(uuid.New())
}

// PlanIDFromString parses a UUID string into a PlanID
func PlanIDFromString(id string) (PlanID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return PlanID{}, xerrors.Errorf("invalid UUID string: %w", err)
	}
	return PlanID(parsedID), nil
}

// String returns the string representation of PlanID
func (id PlanID) String() string {
	return uuid.UUID(id).String()
}

// IsZero checks if PlanID is a zero value
func (id PlanID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

// Scan implements the Scanner interface for database deserialization
func (id *PlanID) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		parsedID, err := uuid.ParseBytes(v)
		if err != nil {
			return err
		}
		*id = PlanID(parsedID)
	case string:
		parsedID, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		*id = PlanID(parsedID)
	default:
		return xerrors.New("invalid UUID type")
	}
	return nil
}

// Value implements the Valuer interface for database serialization
func (id PlanID) Value() (driver.Value, error) {
	return uuid.UUID(id).String(), nil
}
