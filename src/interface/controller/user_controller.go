package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
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
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.CreateUser(ctx.Request.Context(), &in)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func (c *userController) UpdateUser(ctx context.Context, in *userinput.UpdateUserInput) error {
	usecase := userusecase.NewUpdateUser(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.Update(out)
	return nil
}

func (c *userController) UpdateUserHandler(ctx *gin.Context) {
	var in userinput.UpdateUserInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": smperr.BadRequest("Invalid JSON input").Error()})
		return
	}
	in.ID = ctx.Param("id")
	if err := c.UpdateUser(ctx.Request.Context(), &in); err != nil {
		switch e := err.(type) {
		case *smperr.NotFoundErr:
			ctx.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		case *smperr.BadRequestErr:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
