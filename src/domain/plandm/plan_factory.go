package plandm

import (
	"github.com/google/uuid"
)

func GenPlanIfCreate(planID, userID, title, content string, category, consultationFormat, consultationMethod, price, status int) (*Plan, error) {
	parsedPlanID, err := uuid.Parse(planID)
	if err != nil {
		return nil, err
	}

	return newPlan(
		PlanID(parsedPlanID),
		userID,
		title,
		content,
		category,
		consultationFormat,
		consultationMethod,
		price,
		status,
	)
}

func GenWhenRetrieve(planID, userID, title, content string, category, consultationFormat, consultationMethod, price, status int) (*Plan, error) {
	parsedPlanID, err := uuid.Parse(planID)
	if err != nil {
		return nil, err
	}

	return newPlan(
		PlanID(parsedPlanID),
		userID,
		title,
		content,
		category,
		consultationFormat,
		consultationMethod,
		price,
		status,
	)
}
