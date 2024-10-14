package contractrequestdm

import (
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type ContractRequest struct {
	id        ContractRequestID
	planID    plandm.PlanID
	message   string
	createdAt shared.CreatedAt
}

func newContractRequest(id ContractRequestID, planID plandm.PlanID, message string) (*ContractRequest, error) {
	if planID == "" {
		return nil, smperr.BadRequest("planIDは必須です")
	}
	if message == "" {
		return nil, smperr.BadRequest("メッセージは必須です")
	}
	if utf8.RuneCountInString(message) > 500 {
		return nil, smperr.BadRequest("メッセージは500文字以内で入力してください")
	}

	return &ContractRequest{
		id:      id,
		planID:  planID,
		message: message,
	}, nil
}

func (r *ContractRequest) ID() ContractRequestID {
	return r.id
}

func (r *ContractRequest) PlanID() plandm.PlanID {
	return r.planID
}

func (r *ContractRequest) Message() string {
	return r.message
}

func (r *ContractRequest) CreatedAt() shared.CreatedAt {
	return r.createdAt
}
