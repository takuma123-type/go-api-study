package controller

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/planusecase"
	"github.com/takuma123-type/go-api-study/src/usecase/planusecase/planinput"
)

type PlanController struct {
	delivery presenter.PlanPresenter
	planRepo plandm.PlanRepository
}

func NewPlanController(p presenter.PlanPresenter, planRepo plandm.PlanRepository) *PlanController {
	return &PlanController{
		delivery: p,
		planRepo: planRepo,
	}
}

func (c *PlanController) CreatePlan(ctx context.Context, in *planinput.CreatePlanInput) error {
	usecase := planusecase.NewCreatePlan(c.planRepo)
	out, err := usecase.Create(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.CreatePlan(out)
	return nil
}
