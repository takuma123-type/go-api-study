package database

import (
	"context"
	"log"

	"github.com/takuma123-type/go-api-study/src/domain/plandm"
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

func (repo *planRepositoryImpl) UpdateStatus(ctx context.Context, planID string, status uint8) error {
	log.Printf("Updating status for plan ID: %s to status: %d", planID, status)

	query := `
        UPDATE plans
        SET status = ?
        WHERE id = ?`

	if err := repo.db.WithContext(ctx).Exec(query, status, planID).Error; err != nil {
		log.Printf("Failed to update status for plan ID %s: %v", planID, err)
		return err
	}

	return nil
}
