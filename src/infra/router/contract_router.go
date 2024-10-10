package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/infra/rdb"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/contractusecase/contractinput"
)

func NewContractRouter(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.POST("/contract", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			var input contractinput.CreateContractInput
			if err := ctx.ShouldBindJSON(&input); err != nil {
				smperr.HandleError(ctx, err, http.StatusBadRequest)
				return
			}

			contractRepositoryImpl := database.NewContractRepositoryImpl(db)
			planRepositoryImpl := database.NewPlanRepositoryImpl(db)
			err = controller.NewContractController(presenter.NewContractPresenter(ctx), contractRepositoryImpl, planRepositoryImpl).CreateContract(ctx.Request.Context(), &input)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{"message": "Contract created successfully"})
		})
	}
}
