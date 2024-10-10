// src/interface/controller/contract_controller.go
package controller

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/contractdm"
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/contractusecase"
	"github.com/takuma123-type/go-api-study/src/usecase/contractusecase/contractinput"
)

type ContractController struct {
	delivery     presenter.ContractPresenter
	contractRepo contractdm.ContractRepository
	planRepo     plandm.PlanRepository
}

func NewContractController(p presenter.ContractPresenter, contractRepo contractdm.ContractRepository, planRepo plandm.PlanRepository) *ContractController {
	return &ContractController{
		delivery:     p,
		contractRepo: contractRepo,
		planRepo:     planRepo,
	}
}

func (c *ContractController) CreateContract(ctx context.Context, in *contractinput.CreateContractInput) error {
	usecase := contractusecase.NewCreateContract(c.contractRepo, c.planRepo)
	out, err := usecase.Create(ctx, in)
	if err != nil {
		return err
	}

	c.delivery.CreateContract(out)
	return nil
}
