package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
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

func (c *mentorController) FindAllMentorRecruitment(ctx *gin.Context) error {
	usecase := mentorusecase.NewFindAllMentorRecruitment(c.mentorRepo)
	out, err := usecase.Fetch(ctx)
	if err != nil {
		return err
	}
	c.delivery.FindAllMentorRecruitment(out)
	return nil
}
