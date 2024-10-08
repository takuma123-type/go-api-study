package plandm

import "github.com/google/uuid"

type PlanID string

func (id PlanID) String() string {
	return string(id)
}
func NewPlanID() PlanID {
	return PlanID(uuid.New().String())
}
