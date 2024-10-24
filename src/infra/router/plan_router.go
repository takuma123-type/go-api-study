package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/infra/rdb"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/planusecase/planinput"
)

func NewPlanRouter(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.POST("/plan", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			var input planinput.CreatePlanInput
			if err := ctx.ShouldBindJSON(&input); err != nil {
				smperr.HandleError(ctx, err, http.StatusBadRequest)
				return
			}

			planRepositoryImpl := database.NewPlanRepositoryImpl(db)
			err = controller.NewPlanController(presenter.NewPlanPresenter(ctx), planRepositoryImpl).CreatePlan(ctx.Request.Context(), &input)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}
		})

		api.GET("/plan", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			planRepositoryImpl := database.NewPlanRepositoryImpl(db)
			err = controller.NewPlanController(presenter.NewPlanPresenter(ctx), planRepositoryImpl).FindAllPlan(ctx)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}
		})
	}
}
