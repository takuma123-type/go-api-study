package mentordm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/shared"
)

func TestNewMentorRecruitment(t *testing.T) {
	tests := []struct {
		name               string
		userID             string
		title              string
		description        string
		category           int
		consultationFormat int
		consultationMethod int
		budget             int
		period             int
		status             int
		expectErr          bool
	}{
		{
			name:               "Valid mentor recruitment",
			userID:             "123e4567-e89b-12d3-a456-426614174000",
			title:              "Need a mentor for Go",
			description:        "Looking for a mentor to help with Go language.",
			category:           1,
			consultationFormat: 1,
			consultationMethod: 1,
			budget:             1000,
			period:             2,
			status:             1,
			expectErr:          false,
		},
		{
			name:               "Empty title",
			userID:             "123e4567-e89b-12d3-a456-426614174000",
			title:              "",
			description:        "Looking for a mentor to help with Go language.",
			category:           1,
			consultationFormat: 1,
			consultationMethod: 1,
			budget:             1000,
			period:             2,
			status:             1,
			expectErr:          true,
		},
		{
			name:               "Empty description",
			userID:             "123e4567-e89b-12d3-a456-426614174000",
			title:              "Need a mentor for Go",
			description:        "",
			category:           1,
			consultationFormat: 1,
			consultationMethod: 1,
			budget:             1000,
			period:             2,
			status:             1,
			expectErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mentorRecruitmentID := NewMentorRecruitmentID()
			createdAt := shared.NewCreatedAt()

			mentorRecruitment, err := newMentorRecruitment(
				mentorRecruitmentID,
				tt.userID,
				tt.title,
				tt.category,
				tt.consultationFormat,
				tt.consultationMethod,
				tt.budget,
				tt.period,
				tt.status,
				tt.description,
				createdAt,
			)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, mentorRecruitmentID, mentorRecruitment.GetID())
				assert.Equal(t, tt.title, mentorRecruitment.GetTitle())
				assert.Equal(t, tt.description, mentorRecruitment.GetDescription())
			}
		})
	}
}
