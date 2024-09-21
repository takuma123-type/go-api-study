package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/userinput"
)

type userController struct {
	delivery presenter.UserPresenter
	userRepo userdm.UserRepository
}

func NewUserController(p presenter.UserPresenter, userRepo userdm.UserRepository) *userController {
	return &userController{
		delivery: p,
		userRepo: userRepo,
	}
}

func (c *userController) GetUserList(ctx context.Context) error {
	usecase := userusecase.NewGetUserList(c.userRepo)
	out, err := usecase.Exec(ctx)
	if err != nil {
		return err
	}
	c.delivery.UserList(out)
	return nil
}

func (c *userController) GetUserByID(ctx context.Context, in *userinput.GetUserByIDInput) error {
	usecase := userusecase.NewGetUserByID(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.UserByID(out)
	return nil
}

func (c *userController) CreateUser(ctx context.Context, in *userinput.CreateUserInput) error {
	usecase := userusecase.NewCreateUser(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.Create(out)
	return nil
}

func (c *userController) UpdateUser(ctx *gin.Context, in *userinput.UpdateUserInput) error {
	usecase := userusecase.NewUpdateUser(c.userRepo)
	output, err := usecase.Exec(ctx.Request.Context(), in)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, output)
	return nil
}

// gin.HandlerFuncとして使用する場合のラッパー関数
func (c *userController) GetUserByIDHandler(ctx *gin.Context) {
	in := &userinput.GetUserByIDInput{
		ID: ctx.Param("id"),
	}
	err := c.GetUserByID(ctx.Request.Context(), in)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func (c *userController) CreateUserHandler(ctx *gin.Context) {
	var in userinput.CreateUserInput

	// リクエストボディのバインディング
	if err := ctx.ShouldBindJSON(&in); err != nil {
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 正しくバインディングされたかどうかを確認
	log.Printf("Received input: %+v", in)

	// ユーザー作成処理
	err := c.CreateUser(ctx.Request.Context(), &in)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "User created successfully"})
}
