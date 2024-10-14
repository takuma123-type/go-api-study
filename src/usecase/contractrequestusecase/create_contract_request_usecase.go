package contractrequestusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/contractrequestdm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/contractrequestusecase/contractrequestinput"
	"github.com/takuma123-type/go-api-study/src/usecase/contractrequestusecase/contractrequestoutput"
)

type CreateContractRequestUsecase struct {
	contractRequestRepository contractrequestdm.ContractRequestRepository
}

func NewCreateContractRequest(repo contractrequestdm.ContractRequestRepository) *CreateContractRequestUsecase {
	return &CreateContractRequestUsecase{
		contractRequestRepository: repo,
	}
}

func (u *CreateContractRequestUsecase) Create(ctx context.Context, in *contractrequestinput.CreateContractRequestInput) (*contractrequestoutput.CreateContractRequestOutput, error) {
	contractRequest, err := contractrequestdm.GenContractRequestIfCreate(
		in.PlanID,
		in.Message,
	)

	if err != nil {
		return nil, smperr.BadRequest(err.Error())
	}

	if err := u.contractRequestRepository.Store(ctx, contractRequest); err != nil {
		return nil, smperr.Internal(err.Error())
	}

	return &contractrequestoutput.CreateContractRequestOutput{
		ID:      contractRequest.ID().String(),
		PlanID:  contractRequest.PlanID().String(),
		Message: contractRequest.Message(),
	}, nil
}
