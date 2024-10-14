package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/infra/rdb"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/contractapprovalusecase/contractapprovalinput"
)

func NewContractApprovalRouter(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.POST("/contract-approval", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			var input contractapprovalinput.CreateContractApprovalInput
			if err := ctx.ShouldBindJSON(&input); err != nil {
				smperr.HandleError(ctx, err, http.StatusBadRequest)
				return
			}

			contractApprovalRepositoryImpl := database.NewContractApprovalRepositoryImpl(db)
			err = controller.NewContractApprovalController(presenter.NewContractApprovalPresenter(ctx), contractApprovalRepositoryImpl).CreateContractApproval(ctx.Request.Context(), &input)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{"message": "契約承認は正常に送信されました"})
		})
	}
}
