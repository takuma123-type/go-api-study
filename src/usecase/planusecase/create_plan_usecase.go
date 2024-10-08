package planusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/planusecase/planinput"
	"github.com/takuma123-type/go-api-study/src/usecase/planusecase/planoutput"
)

type CreatePlanUsecase struct {
	planRepository plandm.PlanRepository
}

func NewCreatePlan(repo plandm.PlanRepository) *CreatePlanUsecase {
	return &CreatePlanUsecase{
		planRepository: repo,
	}
}

func (p *CreatePlanUsecase) Create(ctx context.Context, in *planinput.CreatePlanInput) (*planoutput.CreatePlanOutput, error) {
	plan, err := plandm.GenPlanIfCreate(
		in.ID,
		in.UserID,
		in.Title,
		in.Content,
		in.Category,
		in.Status,
		in.ConsultationFormat,
		in.Price,
		in.ConsultationMethod,
	)
	if err != nil {
		return nil, smperr.BadRequest(err.Error())
	}

	if err := p.planRepository.Store(ctx, plan); err != nil {
		return nil, smperr.Internal(err.Error())
	}

	return &planoutput.CreatePlanOutput{
		ID:                 plan.ID().String(),
		UserID:             plan.UserID(),
		Title:              plan.Title(),
		Category:           plan.Category(),
		Content:            plan.Content(),
		Status:             plan.Status(),
		ConsultationFormat: plan.ConsultationFormat(),
		Price:              plan.Price(),
		ConsultationMethod: plan.ConsultationMethod(),
	}, nil
}
