package mentordm

import "github.com/takuma123-type/go-api-study/src/domain/shared"

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

func GenWhenRetrieve(
	id MentorRecruitmentID,
	userID string,
	title string,
	category int,
	consultationFormat int,
	consultationMethod int,
	description string,
	budget int,
	period int,
	status int,
	createdAt shared.CreatedAt,
) *MentorRecruitment {
	return &MentorRecruitment{
		id:                 id,
		userID:             userID,
		title:              title,
		category:           category,
		consultationFormat: consultationFormat,
		consultationMethod: consultationMethod,
		description:        description,
		budget:             budget,
		period:             period,
		status:             status,
		createdAt:          createdAt,
	}
}
