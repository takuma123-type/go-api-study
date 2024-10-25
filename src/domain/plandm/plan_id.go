package plandm

import (
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
