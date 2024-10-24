package planinput

type CreatePlanInput struct {
	ID                 string `json:"id,omitempty"`
	UserID             string `json:"user_id,omitempty"`
	Title              string `json:"title,omitempty"`
	Content            string `json:"content,omitempty"`
	Category           uint16 `json:"category,omitempty"`
	Status             uint16 `json:"status,omitempty"`
	ConsultationFormat uint16 `json:"consultation_format,omitempty"`
	Price              uint16 `json:"price,omitempty"`
	ConsultationMethod uint8  `json:"consultation_method,omitempty"`
}
