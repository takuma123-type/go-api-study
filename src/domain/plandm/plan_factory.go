package plandm

func GenPlanIfCreate(id, userID, title, content string, category, consultationFormat, consultationMethod, price, status int) (*Plan, error) {
	planID := PlanID(id)
	return newPlan(
		planID,
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
