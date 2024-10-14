package contractapprovaldm

import "github.com/google/uuid"

type ContractApprovalID string

func (id ContractApprovalID) String() string {
	return string(id)
}

func NewContractApprovalID() ContractApprovalID {
	return ContractApprovalID(uuid.New().String())
}
