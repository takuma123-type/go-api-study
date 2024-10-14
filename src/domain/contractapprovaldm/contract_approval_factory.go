package contractapprovaldm

import (
	"github.com/google/uuid"
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
)

func GenContractApprovalIfCreate(planID, message string) (*ContractApproval, error) {
	contractApprovalID := ContractApprovalID(uuid.New().String())
	return newContractApproval(
		contractApprovalID,
		plandm.PlanID(planID),
		message,
	)
}
