package mentorinput

type FindAllMentorRecruitmentInput struct {
	UserID             string `json:"user_id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Category           int    `json:"category"`
	ConsultationFormat int    `json:"consultation_format"`
	ConsultationMethod int    `json:"consultation_method"`
	Budget             int    `json:"budget"`
	Period             int    `json:"period"`
	Status             int    `json:"status"`
}
