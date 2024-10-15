package presenter

import (
	"net/http"

	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentoroutput"
)

type mentorPresent struct {
	delivery Presenter
}

func NewMentorPresenter(p Presenter) MentorPresenter {
	return &mentorPresent{
		delivery: p,
	}
}

type MentorPresenter interface {
	CreateMentorRecruitment(out *mentoroutput.CreateMentorRecruitmentOutput)
	FindAllMentorRecruitment(out []*mentoroutput.FindAllMentorRecruitmentOutput)
}

func (p *mentorPresent) CreateMentorRecruitment(out *mentoroutput.CreateMentorRecruitmentOutput) {
	p.delivery.JSON(http.StatusCreated, out)
}

func (p *mentorPresent) FindAllMentorRecruitment(out []*mentoroutput.FindAllMentorRecruitmentOutput) {
	p.delivery.JSON(http.StatusOK, out)
}
