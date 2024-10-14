package database

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/contractapprovaldm"
	"gorm.io/gorm"
)

type contractApprovalRepositoryImpl struct {
	db *gorm.DB
}

func NewContractApprovalRepositoryImpl(db *gorm.DB) *contractApprovalRepositoryImpl {
	return &contractApprovalRepositoryImpl{
		db: db,
	}
}

func (repo *contractApprovalRepositoryImpl) Store(ctx context.Context, contractApproval *contractapprovaldm.ContractApproval) error {

	query := `
    INSERT INTO contract_approvals
    (id, plan_id, message)
    VALUES (?, ?, ?)`

	if err := repo.db.WithContext(ctx).Exec(query,
		contractApproval.ID().String(),
		contractApproval.PlanID().String(),
		contractApproval.Message(),
	).Error; err != nil {
		return err
	}
	return nil
}
