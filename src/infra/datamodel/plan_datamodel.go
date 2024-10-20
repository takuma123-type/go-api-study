package datamodel

import (
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

// 構造体名を Plan に変更して GORM が plans テーブルを参照するようにする
type Plan struct {
	ID                 plandm.PlanID `gorm:"primaryKey"`
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

// ToEntity converts PlanModel to a Plan domain entity.
func (m *Plan) ToEntity() *plandm.Plan {
	planID, _ := plandm.PlanIDFromString(m.ID.String())
	plan, _ := plandm.GenPlanIfCreate(
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
