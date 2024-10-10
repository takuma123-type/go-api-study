package contractdm

import "github.com/google/uuid"

type ContractID string

func (id ContractID) String() string {
	return string(id)
}

func NewContractID() ContractID {
	return ContractID(uuid.New().String())
}
