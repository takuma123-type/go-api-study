package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/infra/rdb"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/contractrequestusecase/contractrequestinput"
)

func NewContractRequestRouter(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.POST("/contract-request", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			var input contractrequestinput.CreateContractRequestInput
			if err := ctx.ShouldBindJSON(&input); err != nil {
				smperr.HandleError(ctx, err, http.StatusBadRequest)
				return
			}

			contractRequestRepositoryImpl := database.NewContractRequestRepositoryImpl(db)
			err = controller.NewContractRequestController(presenter.NewContractRequestPresenter(ctx), contractRequestRepositoryImpl).CreateContractRequest(ctx.Request.Context(), &input)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{"message": "契約リクエストは正常に送信されました"})
		})
	}
}
