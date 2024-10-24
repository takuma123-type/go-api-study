package plandm

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type PlanID uuid.UUID

func NewPlanID() PlanID {
	return PlanID(uuid.New())
}

func PlanIDFromString(id string) (PlanID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return PlanID{}, xerrors.Errorf("invalid UUID string: %w", err)
	}
	return PlanID(parsedID), nil
}

func (id PlanID) String() string {
	return uuid.UUID(id).String()
}

func (id PlanID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

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

func (id PlanID) Value() (driver.Value, error) {
	return uuid.UUID(id).String(), nil
}
