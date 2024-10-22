package mentorusecase

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/usecase/mentorusecase/mentoroutput"
)

type FindAllMentorRecruitment struct {
	mentorRepo mentordm.MentorRecruitmentRepository
}

func NewFindAllMentorRecruitment(repo mentordm.MentorRecruitmentRepository) *FindAllMentorRecruitment {
	return &FindAllMentorRecruitment{
		mentorRepo: repo,
	}
}

func (uc *FindAllMentorRecruitment) Fetch(ctx context.Context) ([]*mentoroutput.FindAllMentorRecruitmentOutput, error) {
	mentorRecruitments, err := uc.mentorRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*mentoroutput.FindAllMentorRecruitmentOutput
	for _, mentor := range mentorRecruitments {
		output = append(output, &mentoroutput.FindAllMentorRecruitmentOutput{
			ID:                 mentor.ID().String(),
			UserID:             mentor.UserID(),
			Title:              mentor.Title(),
			Category:           mentor.Category(),
			ConsultationFormat: mentor.ConsultationFormat(),
			Description:        mentor.Description(),
			Budget:             mentor.Budget(),
			Period:             mentor.Period(),
			Status:             mentor.Status(),
			CreatedAt:          mentor.CreatedAt().Value(),
		})
	}

	return output, nil
}
