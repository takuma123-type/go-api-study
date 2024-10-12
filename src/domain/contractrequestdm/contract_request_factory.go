package contractrequestdm

import (
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
)

func GenContractRequestIfCreate(id, planID, message string) (*ContractRequest, error) {
	contractRequestID := ContractRequestID(id)
	return newContractRequest(
		contractRequestID,
		plandm.PlanID(planID),
		message,
	)
}
