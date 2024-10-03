package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentorinput"
)

type mentorController struct {
	delivery   presenter.MentorPresenter
	mentorRepo mentordm.MentorRecruitmentRepository
}

func NewMentorController(p presenter.MentorPresenter, mentorRepo mentordm.MentorRecruitmentRepository) *mentorController {
	return &mentorController{
		delivery:   p,
		mentorRepo: mentorRepo,
	}
}

func (c *mentorController) CreateMentorRecruitment(ctx context.Context, in *mentorinput.CreateMentorRecruitmentInput) error {
	usecase := mentorusecase.NewCreateMentorRecruitment(c.mentorRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.CreateMentorRecruitment(out)
	return nil
}

func (c *mentorController) CreateMentorRecruitmentHandler(ctx *gin.Context) {
	var in mentorinput.CreateMentorRecruitmentInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": smperr.BadRequest("Invalid JSON input").Error()})
		return
	}
	err := c.CreateMentorRecruitment(ctx.Request.Context(), &in)
	if err != nil {
		switch e := err.(type) {
		case *smperr.BadRequestErr:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mentor recruitment created successfully"})
}
