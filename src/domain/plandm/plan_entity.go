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
	content            string
	category           uint16
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

func newPlan(id PlanID, userID string, title string, content string, category uint16, status uint16, consultationFormat uint16, price uint16, consultationMethod uint8) (*Plan, error) {
	if title == "" {
		return nil, smperr.BadRequest("タイトルは空にできません")
	}
	if content == "" {
		return nil, smperr.BadRequest("コンテンツは空にできません")
	}

	if l := utf8.RuneCountInString(title); l > titleLength {
		return nil, smperr.BadRequest("タイトルは%d文字以内にしてください")
	}
	if l := utf8.RuneCountInString(content); l > contentLength {
		return nil, smperr.BadRequest("コンテンツは%d文字以内にしてください")
	}

	return &Plan{
		id:                 id,
		userID:             userID,
		title:              title,
		content:            content,
		category:           category,
		status:             status,
		consultationFormat: consultationFormat,
		price:              price,
		consultationMethod: consultationMethod,
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
