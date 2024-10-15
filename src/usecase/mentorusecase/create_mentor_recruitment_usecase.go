package mentorusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentorinput"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentoroutput"
)

type CreateMentorRecruitmentUsecase struct {
	mentorRecruitmentRepository mentordm.MentorRecruitmentRepository
}

func NewCreateMentorRecruitment(repo mentordm.MentorRecruitmentRepository) *CreateMentorRecruitmentUsecase {
	return &CreateMentorRecruitmentUsecase{
		mentorRecruitmentRepository: repo,
	}
}

func (use *CreateMentorRecruitmentUsecase) Exec(ctx context.Context, in *mentorinput.CreateMentorRecruitmentInput) (*mentoroutput.CreateMentorRecruitmentOutput, error) {
	mentorRecruitment, err := mentordm.GenWhenCreate(
		in.UserID,
		in.Title,
		in.Description,
		in.Category,
		in.ConsultationFormat,
		in.ConsultationMethod,
		in.Budget,
		in.Period,
		in.Status,
	)
	if err != nil {
		return nil, smperr.BadRequest("Invalid input data for creating mentor recruitment")
	}

	if err := use.mentorRecruitmentRepository.Store(ctx, mentorRecruitment); err != nil {
		return nil, smperr.Internal("Store mentor recruitment")
	}

	return &mentoroutput.CreateMentorRecruitmentOutput{
		ID:                 mentorRecruitment.GetID().String(),
		Title:              mentorRecruitment.GetDescription(),
		Description:        mentorRecruitment.GetDescription(),
		CreatedAt:          mentorRecruitment.GetCreatedAt().Time(),
		ConsultationMethod: mentorRecruitment.GetBudget(),
	}, nil
}
