package mentordm

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
