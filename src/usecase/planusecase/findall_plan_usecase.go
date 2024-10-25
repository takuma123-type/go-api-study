package planusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/plandm"
	"github.com/takuma123-type/go-api-study/src/usecase/planusecase/planoutput"
)

type FindAllPlanUsecase struct {
	planRepo plandm.PlanRepository
}

func NewFindAllPlanUsecase(repo plandm.PlanRepository) *FindAllPlanUsecase {
	return &FindAllPlanUsecase{
		planRepo: repo,
	}
}

func (uc *FindAllPlanUsecase) Fetch(ctx context.Context) ([]*planoutput.FindAllPlanOutput, error) {
	plans, err := uc.planRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*planoutput.FindAllPlanOutput, len(plans))
	for i, plan := range plans {
		output[i] = &planoutput.FindAllPlanOutput{
			ID:                 plan.ID().String(),
			UserID:             plan.UserID(),
			Title:              plan.Title(),
			Category:           plan.Category(),
			Content:            plan.Content(),
			Status:             plan.Status(),
			ConsultationFormat: plan.ConsultationFormat(),
			Price:              plan.Price(),
			ConsultationMethod: plan.ConsultationMethod(),
		}
	}

	return output, nil
}
