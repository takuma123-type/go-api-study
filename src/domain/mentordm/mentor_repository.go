package mentordm

import (
	"context"
)

type MentorRecruitmentRepository interface {
	Store(ctx context.Context, mentorRecruitment *MentorRecruitment) error
	FindAll(ctx context.Context) ([]*MentorRecruitment, error)
}
