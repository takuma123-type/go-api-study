package plandm

import (
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type Plan struct {
	id                 PlanID
	userID             string
	title              string
	category           uint16
	content            string
	status             uint16
	consultationFormat uint16
	price              uint16
	consultationMethod uint8
	createdAt          shared.CreatedAt
	updatedAt          shared.UpdatedAt
}

var (
	titleLength   = 255
	contentLength = 2000
)

func newPlan(id PlanID, userID, title, content string, category, status, consultationFormat, price, consultationMethod int) (*Plan, error) {
	if title == "" {
		return nil, smperr.BadRequest("title must not be empty")
	}
	if content == "" {
		return nil, smperr.BadRequest("content must not be empty")
	}

	if l := utf8.RuneCountInString(title); l > titleLength {
		return nil, smperr.BadRequestf("title must be less than %d characters", titleLength)
	}
	if l := utf8.RuneCountInString(content); l > contentLength {
		return nil, smperr.BadRequestf("content must be less than %d characters", contentLength)
	}

	return &Plan{
		id:                 id,
		userID:             userID,
		title:              title,
		category:           uint16(category),
		content:            content,
		status:             uint16(status),
		consultationFormat: uint16(consultationFormat),
		price:              uint16(price),
		consultationMethod: uint8(consultationMethod),
	}, nil
}

func (p *Plan) ID() PlanID {
	return p.id
}

func (p *Plan) UserID() string {
	return p.userID
}

func (p *Plan) Title() string {
	return p.title
}

func (p *Plan) Category() uint16 {
	return p.category
}

func (p *Plan) Content() string {
	return p.content
}

func (p *Plan) Status() uint16 {
	return p.status
}

func (p *Plan) ConsultationFormat() uint16 {
	return p.consultationFormat
}

func (p *Plan) Price() uint16 {
	return p.price
}

func (p *Plan) ConsultationMethod() uint8 {
	return p.consultationMethod
}

func (p *Plan) CreatedAt() shared.CreatedAt {
	return p.createdAt
}

func (p *Plan) UpdatedAt() shared.UpdatedAt {
	return p.updatedAt
}
