package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/userinput"
	"gorm.io/gorm"
)

func NewUserRouter(g *gin.Engine) {
	// DB接続の宣言と初期化部分は不要
	// DSNの設定はミドルウェアなどで扱うべきです

	api := g.Group("/api")
	{
		api.GET("/users", func(ctx *gin.Context) {
			// コンテキストからトランザクション対応のDBを取得
			db := ctx.MustGet("db").(*gorm.DB)
			userRepoImpl := database.NewUserRepositoryImpl(db)
			err := controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).GetUserList(ctx)
			if err != nil {
				ctx.Error(err)
				return
			}
		})

		api.GET("/users/:id", func(ctx *gin.Context) {
			type reqStruct struct {
				ID string `uri:"id"`
			}
			var req reqStruct
			if err := ctx.ShouldBindUri(&req); err != nil {
				ctx.JSON(400, gin.H{"status": "bad request"})
				return
			}

			in := userinput.GetUserByIDInput{ID: req.ID}
			db := ctx.MustGet("db").(*gorm.DB)
			userRepoImpl := database.NewUserRepositoryImpl(db)
			err := controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).GetUserByID(ctx, &in)
			if err != nil {
				ctx.Error(err)
				return
			}
		})

		api.POST("/user", func(ctx *gin.Context) {
			var in userinput.CreateUserInput
			if err := ctx.ShouldBindJSON(&in); err != nil {
				ctx.JSON(400, gin.H{"status": "bad request"})
				return
			}

			db := ctx.MustGet("db").(*gorm.DB)
			userRepoImpl := database.NewUserRepositoryImpl(db)
			err := controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).CreateUser(ctx, &in)
			if err != nil {
				ctx.Error(err)
				return
			}
		})
	}
}
