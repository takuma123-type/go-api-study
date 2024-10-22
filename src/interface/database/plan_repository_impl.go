package database

import (
	"context"
	"log"

	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/infra/datamodel"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"gorm.io/gorm"
)

type planRepositoryImpl struct {
	db *gorm.DB
}

func NewPlanRepositoryImpl(db *gorm.DB) *planRepositoryImpl {
	return &planRepositoryImpl{
		db: db,
	}
}

func (repo *planRepositoryImpl) FindAll(ctx context.Context) ([]*plandm.Plan, error) {
	var planModels []datamodel.Plan
	if err := repo.db.WithContext(ctx).Find(&planModels).Error; err != nil {
		return nil, err
	}

	plans := make([]*plandm.Plan, len(planModels))
	for i, planModel := range planModels {
		plan := planModel.ToEntity()
		plans[i] = plan
	}

	return plans, nil
}

func (repo *planRepositoryImpl) Store(ctx context.Context, plan *plandm.Plan) error {
	log.Printf("Storing plan: %+v", plan)

	query := `
    INSERT INTO plans
    (id, user_id, title, content, category, status, consultation_format, price, consultation_method)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	if err := repo.db.WithContext(ctx).Exec(query,
		plan.ID().String(), plan.UserID(), plan.Title(), plan.Content(), plan.Category(),
		plan.Status(), plan.ConsultationFormat(), plan.Price(), plan.ConsultationMethod(),
	).Error; err != nil {
		log.Printf("Failed to store plan: %v", err)
		return smperr.Internal("store plan")
	}
	return nil
}
