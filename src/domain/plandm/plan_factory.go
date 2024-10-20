package plandm

import (
	"context"

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

func GenWhenFindAll(ctx context.Context, repo PlanRepository) ([]*Plan, error) {
	return repo.FindAll(ctx)
}
