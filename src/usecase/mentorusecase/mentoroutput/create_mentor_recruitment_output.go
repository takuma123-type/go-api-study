package mentoroutput

import "time"

type CreateMentorRecruitmentOutput struct {
	ID                 string    `json:"id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	CreatedAt          time.Time `json:"created_at"`
	ConsultationMethod int       `json:"consultation_method"`
}
