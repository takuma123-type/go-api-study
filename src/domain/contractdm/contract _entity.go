package contractdm

import (
	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type Contract struct {
	id                 ContractID
	planID             string
	userID             string
	contractApprovalID string
	message            string
	status             uint8
	createdAt          shared.CreatedAt
	updatedAt          shared.UpdatedAt
}

func newContract(id ContractID, planID, userID, contractApprovalID, message string, status uint8) (*Contract, error) {
	if planID == "" {
		return nil, smperr.BadRequest("planID must not be empty")
	}
	if userID == "" {
		return nil, smperr.BadRequest("userID must not be empty")
	}
	if contractApprovalID == "" {
		return nil, smperr.BadRequest("contractApprovalID must not be empty")
	}
	if status == 0 {
		return nil, smperr.BadRequest("status must not be 0")
	}
	if message == "" {
		return nil, smperr.BadRequest("message must not be empty")
	}

	return &Contract{
		id:                 id,
		planID:             planID,
		userID:             userID,
		contractApprovalID: contractApprovalID,
		message:            message,
		status:             status,
	}, nil
}

func (c *Contract) ID() ContractID {
	return c.id
}

func (c *Contract) PlanID() string {
	return c.planID
}

func (c *Contract) UserID() string {
	return c.userID
}

func (c *Contract) ContractApprovalID() string {
	return c.contractApprovalID
}

func (c *Contract) Message() string {
	return c.message
}

func (c *Contract) Status() uint8 {
	return c.status
}
