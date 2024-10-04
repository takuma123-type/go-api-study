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
	plan, err := plandm.GenPlanCreate(
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
		ID:                 plan.GetID().String(),
		UserID:             plan.GetUserID(),
		Title:              plan.GetTitle(),
		Category:           plan.GetCategory(),
		Content:            plan.GetContent(),
		Status:             plan.GetStatus(),
		ConsultationFormat: plan.GetConsultationFormat(),
		Price:              plan.GetPrice(),
		ConsultationMethod: plan.GetConsultationMethod(),
	}, nil
}