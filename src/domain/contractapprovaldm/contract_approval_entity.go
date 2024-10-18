package contractapprovaldm

import (
	"unicode/utf8"

	"github.com/takuma123-type/go-api-study/src/domain/contractrequestdm" // ContractRequest に対応するパッケージ
	"github.com/takuma123-type/go-api-study/src/domain/shared"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

type ContractApproval struct {
	id                ContractApprovalID
	contractRequestID contractrequestdm.ContractRequestID // 修正: ContractRequestID に変更
	message           string
	createdAt         shared.CreatedAt
}

func newContractApproval(id ContractApprovalID, contractRequestID contractrequestdm.ContractRequestID, message string) (*ContractApproval, error) { // 修正: ContractRequestID に変更
	if contractRequestID == "" {
		return nil, smperr.BadRequest("contractRequestIDは必須です")
	}
	if message == "" {
		return nil, smperr.BadRequest("メッセージは必須です")
	}
	if utf8.RuneCountInString(message) > 500 {
		return nil, smperr.BadRequest("メッセージは500文字以内で入力してください")
	}

	return &ContractApproval{
		id:                id,
		contractRequestID: contractRequestID, // 修正: ContractRequestID に変更
		message:           message,
	}, nil
}

func (a *ContractApproval) ID() ContractApprovalID {
	return a.id
}

func (a *ContractApproval) ContractRequestID() contractrequestdm.ContractRequestID { // 修正: ContractRequestID に変更
	return a.contractRequestID
}

func (a *ContractApproval) Message() string {
	return a.message
}

func (a *ContractApproval) CreatedAt() shared.CreatedAt {
	return a.createdAt
}
