package database

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/contractrequestdm"
	"gorm.io/gorm"
)

type contractRequestRepositoryImpl struct {
	db *gorm.DB
}

func NewContractRequestRepositoryImpl(db *gorm.DB) *contractRequestRepositoryImpl {
	return &contractRequestRepositoryImpl{
		db: db,
	}
}

func (repo *contractRequestRepositoryImpl) Store(ctx context.Context, contractRequest *contractrequestdm.ContractRequest) error {

	query := `
    INSERT INTO contract_approvals
    (id, plan_id, message)
    VALUES (?, ?, ?)`

	if err := repo.db.WithContext(ctx).Exec(query,
		contractRequest.ID().String(),
		contractRequest.PlanID().String(),
		contractRequest.Message(),
	).Error; err != nil {
		return err
	}
	return nil
}
