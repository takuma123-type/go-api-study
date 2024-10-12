package contractrequestdm

import "github.com/google/uuid"

type ContractRequestID string

func (id ContractRequestID) String() string {
	return string(id)
}

func NewContractRequestID() ContractRequestID {
	return ContractRequestID(uuid.New().String())
}
