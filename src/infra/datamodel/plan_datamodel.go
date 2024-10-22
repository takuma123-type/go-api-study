package datamodel

import (
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

type Plan struct {
	ID                 plandm.PlanID
	UserID             string
	Title              string
	Category           uint16
	Content            string
	Status             uint16
	ConsultationFormat uint16
	Price              uint16
	ConsultationMethod uint8
	CreatedAt          shared.CreatedAt
	UpdatedAt          shared.UpdatedAt
}

func (m *Plan) ToEntity() *plandm.Plan {
	planID, _ := plandm.PlanIDFromString(m.ID.String())
	plan, _ := plandm.GenWhenRetrieve(
		planID.String(),
		m.UserID,
		m.Title,
		m.Content,
		int(m.Category),
		int(m.Status),
		int(m.ConsultationFormat),
		int(m.Price),
		int(m.ConsultationMethod),
	)
	return plan
}
