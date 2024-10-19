package database

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"gorm.io/gorm"
)

type mentorRecruitmentRepositoryImpl struct {
	db *gorm.DB
}

func NewMentorRecruitmentRepositoryImpl(db *gorm.DB) *mentorRecruitmentRepositoryImpl {
	return &mentorRecruitmentRepositoryImpl{
		db: db,
	}
}

func (repo *mentorRecruitmentRepositoryImpl) FindAll(ctx context.Context) ([]*mentordm.MentorRecruitment, error) {
	var mentorRecruitments []*mentordm.MentorRecruitment
	if err := repo.db.WithContext(ctx).Find(&mentorRecruitments).Error; err != nil {
		return nil, err
	}
	return mentorRecruitments, nil
}

func (repo *mentorRecruitmentRepositoryImpl) Store(ctx context.Context, mentorRecruitment *mentordm.MentorRecruitment) error {
	if err := repo.db.WithContext(ctx).Exec(`
        INSERT INTO mentor_recruitments 
        (id, user_id, title, category, consultation_format, consultation_method, description, budget, period, status, created_at) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		mentorRecruitment.GetID().String(), mentorRecruitment.GetUserID(), mentorRecruitment.GetTitle(), mentorRecruitment.GetCategory(),
		mentorRecruitment.GetConsultationFormat(), mentorRecruitment.GetConsultationMethod(), mentorRecruitment.GetDescription(),
		mentorRecruitment.GetBudget(), mentorRecruitment.GetPeriod(), mentorRecruitment.GetStatus(), mentorRecruitment.CreatedAt.Value(),
	).Error; err != nil {
		return err
	}
	return nil
}
