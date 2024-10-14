package contractrequestdm

import (
	"github.com/google/uuid"
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
)

func GenContractRequestIfCreate(planID, message string) (*ContractRequest, error) {
	contractRequestID := ContractRequestID(uuid.New().String())
	return newContractRequest(
		contractRequestID,
		plandm.PlanID(planID),
		message,
	)
}
