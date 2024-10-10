package database

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/takuma123-type/go-api-study/src/domain/contractdm"
	"gorm.io/gorm"
)

type contractRepositoryImpl struct {
	db *gorm.DB
}

func NewContractRepositoryImpl(db *gorm.DB) *contractRepositoryImpl {
	return &contractRepositoryImpl{
		db: db,
	}
}

func (repo *contractRepositoryImpl) Store(ctx context.Context, contract *contractdm.Contract) error {
	log.Printf("Storing contract: %+v", contract)

	contractApprovalID := uuid.New().String()
	approvalQuery := `
        INSERT INTO contract_approvals
        (id, message, plan_id)
        VALUES (?, ?, ?)`

	if err := repo.db.WithContext(ctx).Exec(approvalQuery,
		contractApprovalID,
		"Contract approval pending",
		contract.PlanID(),
	).Error; err != nil {
		log.Printf("Failed to store contract approval: %v", err)
		return err
	}

	contractQuery := `
        INSERT INTO contracts
        (id, user_id, contract_approval_id, message)
        VALUES (?, ?, ?, ?)`

	if err := repo.db.WithContext(ctx).Exec(contractQuery,
		contract.ID().String(),
		contract.UserID(),
		contractApprovalID,
		contract.Message(),
	).Error; err != nil {
		log.Printf("Failed to store contract: %v", err)
		return err
	}

	return nil
}
