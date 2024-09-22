package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/infra/rdb"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/userinput"
)

func NewUserRouter(g *gin.Engine) {
	api := g.Group("/api")
	{
		api.GET("/users", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				ctx.JSON(500, gin.H{"message": (&smperr.DatabaseConnectionError{Reason: err.Error()}).Error()})
				return
			}

			userRepoImpl := database.NewUserRepositoryImpl(db)
			err = controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).GetUserList(ctx)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
		})

		api.GET("/users/:id", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				ctx.JSON(500, gin.H{"message": (&smperr.DatabaseConnectionError{Reason: err.Error()}).Error()})
				return
			}

			type reqStruct struct {
				ID string `uri:"id"`
			}
			var req reqStruct
			if err := ctx.ShouldBindUri(&req); err != nil {
				ctx.JSON(400, gin.H{"message": (&smperr.InvalidURIError{Detail: err.Error()}).Error()})
				return
			}

			in := userinput.GetUserByIDInput{ID: req.ID}
			userRepoImpl := database.NewUserRepositoryImpl(db)
			err = controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).GetUserByID(ctx, &in)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
		})

		api.POST("/user", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				ctx.JSON(500, gin.H{"message": (&smperr.DatabaseConnectionError{Reason: err.Error()}).Error()})
				return
			}

			var in userinput.CreateUserInput
			if err := ctx.ShouldBindJSON(&in); err != nil {
				ctx.JSON(400, gin.H{"message": (&smperr.InvalidJSONError{Detail: err.Error()}).Error()})
				return
			}

			userRepoImpl := database.NewUserRepositoryImpl(db)
			err = controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).CreateUser(ctx, &in)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
		})

		api.PUT("/user/:id", func(ctx *gin.Context) {
			db, err := rdb.GetDBFromContext(ctx)
			if err != nil {
				ctx.JSON(500, gin.H{"message": (&smperr.DatabaseConnectionError{Reason: err.Error()}).Error()})
				return
			}
			id := ctx.Param("id")

			var in userinput.UpdateUserInput
			if err := ctx.ShouldBindJSON(&in); err != nil {
				ctx.JSON(400, gin.H{"message": (&smperr.InvalidJSONError{Detail: err.Error()}).Error()})
				return
			}

			in.ID = id

			userRepoImpl := database.NewUserRepositoryImpl(db)
			err = controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).UpdateUser(ctx, &in)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
		})
	}
}
