package database

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/infra/datamodel"
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
	var mentorRecruitmentModels []datamodel.MentorRecruitmentModel
	if err := repo.db.WithContext(ctx).Find(&mentorRecruitmentModels).Error; err != nil {
		return nil, err
	}

	var mentorRecruitments []*mentordm.MentorRecruitment
	for _, model := range mentorRecruitmentModels {
		mentorRecruitments = append(mentorRecruitments, model.ToEntity())
	}

	return mentorRecruitments, nil
}

func (repo *mentorRecruitmentRepositoryImpl) Store(ctx context.Context, mentorRecruitment *mentordm.MentorRecruitment) error {
	if err := repo.db.WithContext(ctx).Exec(`
        INSERT INTO mentor_recruitments 
        (id, user_id, title, category, consultation_format, consultation_method, description, budget, period, status) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		mentorRecruitment.ID().String(), mentorRecruitment.UserID(), mentorRecruitment.Title(), mentorRecruitment.Category(),
		mentorRecruitment.ConsultationFormat(), mentorRecruitment.ConsultationMethod(), mentorRecruitment.Description(),
		mentorRecruitment.Budget(), mentorRecruitment.Period(), mentorRecruitment.Status(),
	).Error; err != nil {
		return err
	}
	return nil
}
