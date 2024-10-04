package plandm

import (
	"fmt"
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type Plan struct {
	ID                 PlanID `json:"id"`
	UserID             string `json:"user_id,omitempty"`
	Title              string `json:"title,omitempty"`
	Category           int    `json:"category,omitempty"`
	Content            string `json:"content,omitempty"`
	Status             int    `json:"status,omitempty"`
	ConsultationFormat int    `json:"consultation_format,omitempty"`
	Price              int    `json:"price,omitempty"`
	ConsultationMethod int    `json:"consultation_method,omitempty"`
	CreatedAt          shared.CreatedAt
	UpdatedAt          shared.UpdatedAt
}

var (
	titleLength   = 255
	contentLength = 2000
)

func newPlan(userID, title, content string, category, status, consultationFormat, price, consultationMethod int) (*Plan, error) {
	if title == "" {
		return nil, smperr.BadRequest("title must not be empty")
	}
	if content == "" {
		return nil, smperr.BadRequest("content must not be empty")
	}

	if l := utf8.RuneCountInString(title); l > titleLength {
		return nil, smperr.BadRequest(
			fmt.Sprintf("title must be less than %d characters", titleLength),
		)
	}
	if l := utf8.RuneCountInString(content); l > contentLength {
		return nil, smperr.BadRequest(
			fmt.Sprintf("content must be less than %d characters", contentLength),
		)
	}

	return &Plan{
		ID:                 NewPlanID(),
		UserID:             userID,
		Title:              title,
		Category:           category,
		Content:            content,
		Status:             status,
		ConsultationFormat: consultationFormat,
		Price:              price,
		ConsultationMethod: consultationMethod,
	}, nil
}

func (p *Plan) GetID() PlanID {
	return p.ID
}

func (p *Plan) GetUserID() string {
	return p.UserID
}

func (p *Plan) GetTitle() string {
	return p.Title
}

func (p *Plan) GetCategory() int {
	return p.Category
}

func (p *Plan) GetContent() string {
	return p.Content
}

func (p *Plan) GetStatus() int {
	return p.Status
}

func (p *Plan) GetConsultationFormat() int {
	return p.ConsultationFormat
}

func (p *Plan) GetPrice() int {
	return p.Price
}

func (p *Plan) GetConsultationMethod() int {
	return p.ConsultationMethod
}

func (p *Plan) GetCreatedAt() shared.CreatedAt {
	return p.CreatedAt
}

func (p *Plan) GetUpdatedAt() shared.UpdatedAt {
	return p.UpdatedAt
}
