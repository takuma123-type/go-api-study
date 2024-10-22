package datamodel

import (
	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

type MentorRecruitmentModel struct {
	ID                 mentordm.MentorRecruitmentID `gorm:"primaryKey"`
	UserID             string
	Title              string
	Category           int
	ConsultationFormat int
	ConsultationMethod int
	Description        string
	Budget             int
	Period             int
	Status             int
	CreatedAt          shared.CreatedAt
}

func (MentorRecruitmentModel) TableName() string {
	return "mentor_recruitments"
}

func (m *MentorRecruitmentModel) ToEntity() *mentordm.MentorRecruitment {
	return mentordm.GenWhenRetrieve(
		m.ID,
		m.UserID,
		m.Title,
		m.Category,
		m.ConsultationFormat,
		m.ConsultationMethod,
		m.Description,
		m.Budget,
		m.Period,
		m.Status,
		m.CreatedAt,
	)
}
