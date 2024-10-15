package mentoroutput

import "time"

type FindAllMentorRecruitmentOutput struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	Title              string    `json:"title"`
	Category           int       `json:"category"`
	ConsultationFormat int       `json:"consultation_format"`
	Description        string    `json:"description"`
	Budget             int       `json:"budget"`
	Period             int       `json:"period"`
	Status             int       `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
}
