package plandm

import (
	"github.com/google/uuid"
)

func GenPlanIfCreate(userID string, title string, content string, category uint16, status uint16, consultationFormat uint16, price uint16, consultationMethod uint8) (*Plan, error) {
	newPlanID := uuid.New()

	return newPlan(
		PlanID(newPlanID),
		userID,
		title,
		content,
		category,
		status,
		consultationFormat,
		price,
		consultationMethod,
	)
}

func GenWhenRetrieve(id PlanID, userID string, title string, content string, category uint16, status uint16, consultationFormat uint16, price uint16, consultationMethod uint8) (*Plan, error) {

	return newPlan(
		id,
		userID,
		title,
		content,
		category,
		status,
		consultationFormat,
		price,
		consultationMethod,
	)
}
