package contractrequestdm

import (
	"github.com/google/uuid"
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
)

func GenContractRequestIfCreate(planID, message string) (*ContractRequest, error) {
	parsedPlanID, err := plandm.PlanIDFromString(planID)
	if err != nil {
		return nil, err
	}

	contractRequestID := ContractRequestID(uuid.New().String())
	return newContractRequest(
		contractRequestID,
		parsedPlanID,
		message,
	)
}
