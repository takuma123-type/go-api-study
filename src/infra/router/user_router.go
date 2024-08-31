package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/database"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/userinput"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewUserRouter(g *gin.Engine) {
	dsn := "root:password@tcp(db:3306)/golang-db?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic("failed to connect database")
	}

	api := g.Group("/api")
	{
		api.GET("/users", func(ctx *gin.Context) {
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

			userRepoImpl := database.NewUserRepositoryImpl(db)
			err := controller.NewUserController(presenter.NewUserPresenter(ctx), userRepoImpl).CreateUser(ctx, &in)
			if err != nil {
				ctx.Error(err)
				return
			}
		})
	}
}
