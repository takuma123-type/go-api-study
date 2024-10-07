package planoutput

type CreatePlanOutput struct {
	ID                 string `json:"id,omitempty"`
	UserID             string `json:"user_id,omitempty"`
	Title              string `json:"title,omitempty"`
	Category           uint16 `json:"category,omitempty"`
	Content            string `json:"content,omitempty"`
	Status             uint16 `json:"status,omitempty"`
	ConsultationFormat uint16 `json:"consultation_format,omitempty"`
	Price              uint16 `json:"price,omitempty"`
	ConsultationMethod int    `json:"consultation_method,omitempty"`
}
