package mentordm

import "github.com/takuma123-type/go-api-study/src/domain/shared"

func GenWhenCreate(userID, title, description string, category, consultationFormat, consultationMethod, budget, period, status int) (*MentorRecruitment, error) {
	return newMentorRecruitment(NewMentorRecruitmentID(), userID, title, category, consultationFormat, consultationMethod, budget, period, status, description, shared.NewCreatedAt())
}

func GenForTest(id MentorRecruitmentID, userID, title, description string, category, consultationFormat, consultationMethod, budget, period, status int) (*MentorRecruitment, error) {
	return newMentorRecruitment(id, userID, title, category, consultationFormat, consultationMethod, budget, period, status, description, shared.NewCreatedAt())
}
