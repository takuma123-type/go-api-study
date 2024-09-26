package mentorusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentorinput"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentoroutput"
)

const (
	errInvalidInputData = "Invalid input data for creating mentor recruitment"
	errStoreOperation   = "Store mentor recruitment"
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
		return nil, smperr.BadRequest(errInvalidInputData)
	}

	if err := use.mentorRecruitmentRepository.Store(ctx, mentorRecruitment); err != nil {
		return nil, err
	}

	return &mentoroutput.CreateMentorRecruitmentOutput{
		ID:                 mentorRecruitment.GetID().String(),
		Title:              mentorRecruitment.GetTitle(),
		Description:        mentorRecruitment.GetDescription(),
		CreatedAt:          mentorRecruitment.GetCreatedAt().Value(),
		ConsultationMethod: mentorRecruitment.GetConsultationMethod(),
	}, nil
}
