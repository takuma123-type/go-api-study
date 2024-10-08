package presenter

import (
	"net/http"

	"github.com/takuma123-type/go-api-study/src/usecase/planusecase/planoutput"
)

type planPresent struct {
	delivery Presenter
}

func NewPlanPresenter(p Presenter) PlanPresenter {
	return &planPresent{
		delivery: p,
	}
}

type PlanPresenter interface {
	CreatePlan(out *planoutput.CreatePlanOutput)
}

func (p *planPresent) CreatePlan(out *planoutput.CreatePlanOutput) {
	p.delivery.JSON(http.StatusCreated, out)
}
