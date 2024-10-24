package datamodel

import (
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
)

type Plan struct {
	ID                 string
	UserID             string
	Title              string
	Content            string
	Category           uint16
	Status             uint16
	ConsultationFormat uint16
	Price              uint16
	ConsultationMethod uint8
}

func (m *Plan) ToEntity() (*plandm.Plan, error) {
	return plandm.GenWhenRetrieve(
		m.ID,
		m.UserID,
		m.Title,
		m.Content,
		m.Category,
		m.Status,
		m.ConsultationFormat,
		m.Price,
		m.ConsultationMethod,
	)
}
