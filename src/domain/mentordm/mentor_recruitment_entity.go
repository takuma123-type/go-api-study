package mentordm

import (
	"fmt"
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type MentorRecruitment struct {
	ID                 MentorRecruitmentID
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
	UpdatedAt          shared.CreatedAt
}

var (
	titleLength       = 255
	descriptionLength = 2000
)

func newMentorRecruitment(userID, title, description string, category, consultationFormat, consultationMethod, budget, period, status int) (*MentorRecruitment, error) {
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
		ID:                 NewMentorRecruitmentID(),
		UserID:             userID,
		Title:              title,
		Category:           category,
		ConsultationFormat: consultationFormat,
		ConsultationMethod: consultationMethod,
		Description:        description,
		Budget:             budget,
		Period:             period,
		Status:             status,
		CreatedAt:          shared.NewCreatedAt(),
		UpdatedAt:          shared.NewCreatedAt(),
	}, nil
}

func (m *MentorRecruitment) GetID() MentorRecruitmentID {
	return m.ID
}

func (m *MentorRecruitment) GetUserID() string {
	return m.UserID
}

func (m *MentorRecruitment) GetTitle() string {
	return m.Title
}

func (m *MentorRecruitment) GetCategory() int {
	return m.Category
}

func (m *MentorRecruitment) GetConsultationFormat() int {
	return m.ConsultationFormat
}

func (m *MentorRecruitment) GetConsultationMethod() int {
	return m.ConsultationMethod
}

func (m *MentorRecruitment) GetDescription() string {
	return m.Description
}

func (m *MentorRecruitment) GetBudget() int {
	return m.Budget
}

func (m *MentorRecruitment) GetPeriod() int {
	return m.Period
}

func (m *MentorRecruitment) GetStatus() int {
	return m.Status
}

func (m *MentorRecruitment) GetCreatedAt() shared.CreatedAt {
	return m.CreatedAt
}

func (m *MentorRecruitment) GetUpdatedAt() shared.CreatedAt {
	return m.UpdatedAt
}
