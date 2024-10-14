package contractapprovaldm

import (
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type ContractApproval struct {
	id        ContractApprovalID
	planID    plandm.PlanID
	message   string
	createdAt shared.CreatedAt
}

func newContractApproval(id ContractApprovalID, planID plandm.PlanID, message string) (*ContractApproval, error) {
	if planID == "" {
		return nil, smperr.BadRequest("planIDは必須です")
	}
	if message == "" {
		return nil, smperr.BadRequest("メッセージは必須です")
	}
	if utf8.RuneCountInString(message) > 500 {
		return nil, smperr.BadRequest("メッセージは500文字以内で入力してください")
	}

	return &ContractApproval{
		id:      id,
		planID:  planID,
		message: message,
	}, nil
}

func (a *ContractApproval) ID() ContractApprovalID {
	return a.id
}

func (a *ContractApproval) PlanID() plandm.PlanID {
	return a.planID
}

func (a *ContractApproval) Message() string {
	return a.message
}

func (a *ContractApproval) CreatedAt() shared.CreatedAt {
	return a.createdAt
}
