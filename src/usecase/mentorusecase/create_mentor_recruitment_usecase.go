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
	errMentorNotFound   = "Mentor recruitment could not be found"
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
		if smperr.IsRecordNotFound(err) {
			return nil, smperr.NotFound(errMentorNotFound)
		}
		return nil, &smperr.DatabaseError{
			Operation: errStoreOperation,
			Err:       err,
		}
	}

	return &mentoroutput.CreateMentorRecruitmentOutput{
		ID:                 mentorRecruitment.ID.String(),
		Title:              mentorRecruitment.GetTitle(),
		Description:        mentorRecruitment.GetDescription(),
		CreatedAt:          mentorRecruitment.GetCreatedAt().Value(),
		ConsultationMethod: mentorRecruitment.ConsultationMethod,
	}, nil
}
