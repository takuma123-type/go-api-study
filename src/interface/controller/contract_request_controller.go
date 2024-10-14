package controller

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/contractrequestdm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/contractrequestusecase"
	"github.com/takuma123-type/go-api-study/src/usecase/contractrequestusecase/contractrequestinput"
)

type ContractRequestController struct {
	delivery            presenter.ContractRequestPresenter
	contractRequestRepo contractrequestdm.ContractRequestRepository
}

func NewContractRequestController(p presenter.ContractRequestPresenter, contractRequestRepo contractrequestdm.ContractRequestRepository) *ContractRequestController {
	return &ContractRequestController{
		delivery:            p,
		contractRequestRepo: contractRequestRepo,
	}
}

func (c *ContractRequestController) CreateContractRequest(ctx context.Context, in *contractrequestinput.CreateContractRequestInput) error {
	usecase := contractrequestusecase.NewCreateContractRequest(c.contractRequestRepo)
	out, err := usecase.Create(ctx, in)
	if err != nil {
		return err
	}

	c.delivery.CreateContractRequest(out)
	return nil
}