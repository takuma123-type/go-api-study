package plandm

func GenPlanIfCreate(userID, title, content string, category, consultationFormat, consultationMethod, price, status int) (*Plan, error) {
	return newPlan(
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
