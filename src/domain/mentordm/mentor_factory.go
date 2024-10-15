package mentordm

import (
	"context"
)

func GenWhenCreate(userID, title, description string, category, consultationFormat, consultationMethod, budget, period, status int) (*MentorRecruitment, error) {
	return newMentorRecruitment(
		userID,
		title,
		description,
		category,
		consultationFormat,
		consultationMethod,
		budget,
		period,
		status,
	)
}

func GenWhenFindAll(ctx context.Context, repo MentorRecruitmentRepository) ([]*MentorRecruitment, error) {
	return repo.FindAll(ctx)
}
