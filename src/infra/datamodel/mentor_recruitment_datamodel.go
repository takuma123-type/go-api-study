package database

import (
	"context"
	"log"

	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
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
	log.Printf("Storing mentor recruitment: %+v", mentorRecruitment)

	if err := repo.db.WithContext(ctx).Create(&mentorRecruitment).Error; err != nil {
		log.Printf("Failed to store mentor recruitment: %v", err)
		return smperr.Internal("store mentor recruitment")
	}
	return nil
}
