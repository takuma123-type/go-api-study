package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/infra/rdb"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentorinput"
)

func NewMentorRouter(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.POST("/mentor_recruitment", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}

			var in mentorinput.CreateMentorRecruitmentInput
			if err := ctx.ShouldBindJSON(&in); err != nil {
				smperr.HandleError(ctx, err, http.StatusBadRequest)
				return
			}

			mentorRepoImpl := database.NewMentorRecruitmentRepositoryImpl(db)
			err = controller.NewMentorController(presenter.NewMentorPresenter(ctx), mentorRepoImpl).CreateMentorRecruitment(ctx, &in)
			if err != nil {
				smperr.HandleError(ctx, err, http.StatusInternalServerError)
				return
			}
		})
	}
}
