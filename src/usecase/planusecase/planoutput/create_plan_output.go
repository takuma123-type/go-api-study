package planoutput

type CreatePlanOutput struct {
	ID                 string `json:"id,omitempty"`
	UserID             string `json:"user_id,omitempty"`
	Title              string `json:"title,omitempty"`
	Category           int    `json:"category,omitempty"`
	Content            string `json:"content,omitempty"`
	Status             int    `json:"status,omitempty"`
	ConsultationFormat int    `json:"consultation_format,omitempty"`
	Price              int    `json:"price,omitempty"`
	ConsultationMethod int    `json:"consultation_method,omitempty"`
}
