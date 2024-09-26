package mentordm

import (
	"fmt"
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type MentorRecruitment struct {
	id                 MentorRecruitmentID
	userID             string
	title              string
	category           int
	consultationFormat int
	consultationMethod int
	description        string
	budget             int
	period             int
	status             int
	createdAt          shared.CreatedAt
	updatedAt          shared.CreatedAt
}

var (
	titleLength       = 255
	descriptionLength = 2000
)

func NewMentorRecruitment(id MentorRecruitmentID, userID, title string, category, consultationFormat, consultationMethod, budget, period, status int, description string, createdAt shared.CreatedAt) (*MentorRecruitment, error) {
	return newMentorRecruitment(id, userID, title, category, consultationFormat, consultationMethod, budget, period, status, description, createdAt)
}

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
		id:                 id,
		userID:             userID,
		title:              title,
		category:           category,
		consultationFormat: consultationFormat,
		consultationMethod: consultationMethod,
		description:        description,
		budget:             budget,
		period:             period,
		status:             status,
		createdAt:          createdAt,
		updatedAt:          createdAt,
	}, nil
}

func (m *MentorRecruitment) GetID() MentorRecruitmentID {
	return m.id
}

func (m *MentorRecruitment) GetUserID() string {
	return m.userID
}

func (m *MentorRecruitment) GetTitle() string {
	return m.title
}

func (m *MentorRecruitment) GetCategory() int {
	return m.category
}

func (m *MentorRecruitment) GetConsultationFormat() int {
	return m.consultationFormat
}

func (m *MentorRecruitment) GetConsultationMethod() int {
	return m.consultationMethod
}

func (m *MentorRecruitment) GetDescription() string {
	return m.description
}

func (m *MentorRecruitment) GetBudget() int {
	return m.budget
}

func (m *MentorRecruitment) GetPeriod() int {
	return m.period
}

func (m *MentorRecruitment) GetStatus() int {
	return m.status
}

func (m *MentorRecruitment) GetCreatedAt() shared.CreatedAt {
	return m.createdAt
}

func (m *MentorRecruitment) GetUpdatedAt() shared.CreatedAt {
	return m.updatedAt
}
