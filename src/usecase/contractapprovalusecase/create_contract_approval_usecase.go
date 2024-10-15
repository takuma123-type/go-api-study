package contractapprovalusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/contractapprovaldm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/contractapprovalusecase/contractapprovalinput"
	"github.com/takuma123-type/go-api-study/src/usecase/contractapprovalusecase/contractapprovaloutput"
)

type CreateContractApprovalUsecase struct {
	contractApprovalRepository contractapprovaldm.ContractApprovalRepository
}

func NewCreateContractApproval(repo contractapprovaldm.ContractApprovalRepository) *CreateContractApprovalUsecase {
	return &CreateContractApprovalUsecase{
		contractApprovalRepository: repo,
	}
}

func (u *CreateContractApprovalUsecase) Create(ctx context.Context, in *contractapprovalinput.CreateContractApprovalInput) (*contractapprovaloutput.CreateContractApprovalOutput, error) {
	contractApproval, err := contractapprovaldm.GenContractApprovalIfCreate(
		in.ContractRequestID,
		in.Message,
	)

	if err != nil {
		return nil, smperr.BadRequest(err.Error())
	}

	if err := u.contractApprovalRepository.Store(ctx, contractApproval); err != nil {
		return nil, smperr.Internal(err.Error())
	}

	return &contractapprovaloutput.CreateContractApprovalOutput{
		ID:                contractApproval.ID().String(),
		ContractRequestID: contractApproval.ContractRequestID().String(),
		Message:           contractApproval.Message(),
	}, nil
}
