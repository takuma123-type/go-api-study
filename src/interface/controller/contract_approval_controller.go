package controller

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/contractapprovaldm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/contractapprovalusecase"
	"github.com/takuma123-type/go-api-study/src/usecase/contractapprovalusecase/contractapprovalinput"
)

type contractApprovalController struct {
	delivery             presenter.ContractApprovalPresenter
	contractApprovalRepo contractapprovaldm.ContractApprovalRepository
}

func NewContractApprovalController(p presenter.ContractApprovalPresenter, contractApprovalRepo contractapprovaldm.ContractApprovalRepository) *contractApprovalController {
	return &contractApprovalController{
		delivery:             p,
		contractApprovalRepo: contractApprovalRepo,
	}
}

func (c *contractApprovalController) CreateContractApproval(ctx context.Context, in *contractapprovalinput.CreateContractApprovalInput) error {
	usecase := contractapprovalusecase.NewCreateContractApproval(c.contractApprovalRepo)
	out, err := usecase.Create(ctx, in)
	if err != nil {
		return err
	}

	c.delivery.CreateContractApproval(out)
	return nil
}
