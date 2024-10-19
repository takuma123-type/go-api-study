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
			ID:                 mentor.GetID().String(),
			UserID:             mentor.GetUserID(),
			Title:              mentor.GetTitle(),
			Category:           mentor.GetCategory(),
			ConsultationFormat: mentor.GetConsultationFormat(),
			Description:        mentor.GetDescription(),
			Budget:             mentor.GetBudget(),
			Period:             mentor.GetPeriod(),
			Status:             mentor.GetStatus(),
			CreatedAt:          mentor.GetCreatedAt().Value(),
		})
	}

	return output, nil
}
