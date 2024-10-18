package contractapprovaldm

import (
	"github.com/google/uuid"
	"github.com/takuma123-type/go-api-study/src/domain/contractrequestdm"
)

func GenContractApprovalIfCreate(contractRequestID, message string) (*ContractApproval, error) {
	contractApprovalID := ContractApprovalID(uuid.New().String())
	return newContractApproval(
		contractApprovalID,
		contractrequestdm.ContractRequestID(contractRequestID),
		message,
	)
}
