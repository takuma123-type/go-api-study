package datamodel

import (
	"time"

	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

type MentorRecruitmentDataModel struct {
	ID                 string    `gorm:"column:id;primary_key"`
	UserID             string    `gorm:"column:user_id"`
	Title              string    `gorm:"column:title"`
	Category           int       `gorm:"column:category"`
	ConsultationFormat int       `gorm:"column:consultation_format"`
	ConsultationMethod int       `gorm:"column:consultation_method"`
	Description        string    `gorm:"column:description"`
	Budget             int       `gorm:"column:budget"`
	Period             int       `gorm:"column:period"`
	Status             int       `gorm:"column:status"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}

func (d *MentorRecruitmentDataModel) ToDomain() (*mentordm.MentorRecruitment, error) {
	return mentordm.NewMentorRecruitment(
		mentordm.MentorRecruitmentID(d.ID),
		d.UserID,
		d.Title,
		d.Category,
		d.ConsultationFormat,
		d.ConsultationMethod,
		d.Budget,
		d.Period,
		d.Status,
		d.Description,
		shared.CreatedAt(d.CreatedAt),
	)
}

func FromDomain(m *mentordm.MentorRecruitment) *MentorRecruitmentDataModel {
	return &MentorRecruitmentDataModel{
		ID:                 string(m.GetID()),
		UserID:             m.GetUserID(),
		Title:              m.GetTitle(),
		Category:           m.GetCategory(),
		ConsultationFormat: m.GetConsultationFormat(),
		ConsultationMethod: m.GetConsultationMethod(),
		Description:        m.GetDescription(),
		Budget:             m.GetBudget(),
		Period:             m.GetPeriod(),
		Status:             m.GetStatus(),
		CreatedAt:          time.Time(m.GetCreatedAt()),
		UpdatedAt:          time.Time(m.GetUpdatedAt()),
	}
}
