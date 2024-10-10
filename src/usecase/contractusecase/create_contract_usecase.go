package contractusecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/takuma123-type/go-api-study/src/domain/contractdm"
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/contractusecase/contractinput"
	"github.com/takuma123-type/go-api-study/src/usecase/contractusecase/contractoutput"
)

type CreateContractUsecase struct {
	contractRepository contractdm.ContractRepository
	planRepository     plandm.PlanRepository
}

func NewCreateContract(contractRepo contractdm.ContractRepository, planRepo plandm.PlanRepository) *CreateContractUsecase {
	return &CreateContractUsecase{
		contractRepository: contractRepo,
		planRepository:     planRepo,
	}
}

func (c *CreateContractUsecase) Create(ctx context.Context, in *contractinput.CreateContractInput) (*contractoutput.CreateContractOutput, error) {
	if in.ContractApprovalID == "" {
		in.ContractApprovalID = uuid.New().String()
	}

	contractID := contractdm.NewContractID()

	contract, err := contractdm.GenContractIfCreate(
		contractID.String(),
		in.PlanID,
		in.UserID,
		in.ContractApprovalID,
		in.Message,
		in.Status,
	)
	if err != nil {
		return nil, smperr.BadRequest(err.Error())
	}

	if err := c.contractRepository.Store(ctx, contract); err != nil {
		return nil, smperr.Internal(err.Error())
	}

	if err := c.planRepository.UpdateStatus(ctx, in.PlanID, 2); err != nil {
		return nil, smperr.Internal(err.Error())
	}

	return &contractoutput.CreateContractOutput{
		ID:                 contract.ID().String(),
		PlanID:             contract.PlanID(),
		UserID:             contract.UserID(),
		ContractApprovalID: contract.ContractApprovalID(),
		Message:            contract.Message(),
		Status:             contract.Status(),
	}, nil
}
