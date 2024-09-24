package mentordm

import (
	"fmt"
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type MentorRecruitment struct {
	ID                 MentorRecruitmentID `gorm:"column:id"`
	UserID             string              `gorm:"column:user_id"`
	Title              string              `gorm:"column:title"`
	Category           int                 `gorm:"column:category"`
	ConsultationFormat int                 `gorm:"column:consultation_format"`
	ConsultationMethod int                 `gorm:"column:consultation_method"`
	Description        string              `gorm:"column:description"`
	Budget             int                 `gorm:"column:budget"`
	Period             int                 `gorm:"column:period"`
	Status             int                 `gorm:"column:status"`
	CreatedAt          shared.CreatedAt    `gorm:"column:created_at"`
	UpdatedAt          shared.CreatedAt    `gorm:"column:updated_at"`
}

func (m *MentorRecruitment) GetTitle() string {
	return m.Title
}

func (m *MentorRecruitment) GetDescription() string {
	return m.Description
}

func (m *MentorRecruitment) GetCreatedAt() shared.CreatedAt {
	return m.CreatedAt
}

var (
	titleLength       = 255
	descriptionLength = 2000
)

func newMentorRecruitment(id MentorRecruitmentID, userID, title string, category, consultationFormat, consultationMethod, budget, period, status int, description string, createdAt shared.CreatedAt) (*MentorRecruitment, error) {
	if title == "" {
		return nil, smperr.BadRequest("title must not be empty")
	}
	if description == "" {
		return nil, smperr.BadRequest("description must not be empty")
	}

	if l := utf8.RuneCountInString(title); l > titleLength {
		return nil, smperr.BadRequest(
			fmt.Sprintf("title must be less than %d characters", titleLength),
		)
	}
	if l := utf8.RuneCountInString(description); l > descriptionLength {
		return nil, smperr.BadRequest(
			fmt.Sprintf("description must be less than %d characters", descriptionLength),
		)
	}

	return &MentorRecruitment{
		ID:                 id,
		UserID:             userID,
		Title:              title,
		Category:           category,
		ConsultationFormat: consultationFormat,
		ConsultationMethod: consultationMethod,
		Description:        description,
		Budget:             budget,
		Period:             period,
		Status:             status,
		CreatedAt:          createdAt,
		UpdatedAt:          createdAt,
	}, nil
}

func (m *MentorRecruitment) UpdateMentorRecruitment(title, description string, category, consultationFormat, consultationMethod, budget, period, status int) error {
	if title == "" {
		return smperr.BadRequest("title must not be empty")
	}
	if description == "" {
		return smperr.BadRequest("description must not be empty")
	}

	if l := utf8.RuneCountInString(title); l > titleLength {
		return smperr.BadRequest(
			fmt.Sprintf("title must be less than %d characters", titleLength),
		)
	}
	if l := utf8.RuneCountInString(description); l > descriptionLength {
		return smperr.BadRequest(
			fmt.Sprintf("description must be less than %d characters", descriptionLength),
		)
	}

	m.Title = title
	m.Description = description
	m.Category = category
	m.ConsultationFormat = consultationFormat
	m.ConsultationMethod = consultationMethod
	m.Budget = budget
	m.Period = period
	m.Status = status

	return nil
}
