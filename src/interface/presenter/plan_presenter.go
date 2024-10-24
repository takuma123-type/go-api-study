package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/usecase/planusecase/planoutput"
)

type PlanPresenter interface {
	CreatePlan(output *planoutput.CreatePlanOutput)
	FindAllPlan(output []*planoutput.FindAllPlanOutput)
}

type PlanPresenterImpl struct {
	ctx *gin.Context
}

func NewPlanPresenter(ctx *gin.Context) PlanPresenter {
	return &PlanPresenterImpl{ctx: ctx}
}

func (p *PlanPresenterImpl) CreatePlan(output *planoutput.CreatePlanOutput) {
	p.ctx.JSON(http.StatusOK, output)
}

func (p *PlanPresenterImpl) FindAllPlan(output []*planoutput.FindAllPlanOutput) {
	p.ctx.JSON(http.StatusOK, output)
}
