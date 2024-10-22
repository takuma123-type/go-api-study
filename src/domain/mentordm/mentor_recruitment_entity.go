package mentordm

import (
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

func newMentorRecruitment(userID, title, description string, category, consultationFormat, consultationMethod, budget, period, status int) (*MentorRecruitment, error) {
	if title == "" {
		return nil, smperr.BadRequest("title must not be empty")
	}
	if description == "" {
		return nil, smperr.BadRequest("description must not be empty")
	}

	if l := utf8.RuneCountInString(title); l > titleLength {
		return nil, smperr.BadRequest("title must be less than 255 characters")
	}
	if l := utf8.RuneCountInString(description); l > descriptionLength {
		return nil, smperr.BadRequest("description must be less than 2000 characters")
	}

	return &MentorRecruitment{
		id:                 NewMentorRecruitmentID(),
		userID:             userID,
		title:              title,
		category:           category,
		consultationFormat: consultationFormat,
		consultationMethod: consultationMethod,
		description:        description,
		budget:             budget,
		period:             period,
		status:             status,
	}, nil
}

func (m *MentorRecruitment) ID() MentorRecruitmentID {
	return m.id
}

func (m *MentorRecruitment) UserID() string {
	return m.userID
}

func (m *MentorRecruitment) Title() string {
	return m.title
}

func (m *MentorRecruitment) Category() int {
	return m.category
}

func (m *MentorRecruitment) ConsultationFormat() int {
	return m.consultationFormat
}

func (m *MentorRecruitment) ConsultationMethod() int {
	return m.consultationMethod
}

func (m *MentorRecruitment) Description() string {
	return m.description
}

func (m *MentorRecruitment) Budget() int {
	return m.budget
}

func (m *MentorRecruitment) Period() int {
	return m.period
}

func (m *MentorRecruitment) Status() int {
	return m.status
}

func (m *MentorRecruitment) CreatedAt() shared.CreatedAt {
	return m.createdAt
}

func (m *MentorRecruitment) UpdatedAt() shared.CreatedAt {
	return m.updatedAt
}
