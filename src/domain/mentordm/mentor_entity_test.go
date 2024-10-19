package mentordm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/mentordm"
)

func TestNewMentorRecruitment(t *testing.T) {
	tests := []struct {
		name                 string
		userID               string
		title                string
		description          string
		category             int
		consultationFormat   int
		consultationMethod   int
		budget               int
		period               int
		status               int
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:               "Valid Mentor Recruitment",
			userID:             "user123",
			title:              "Valid Title",
			description:        "Valid Description",
			category:           1,
			consultationFormat: 1,
			consultationMethod: 1,
			budget:             1000,
			period:             30,
			status:             1,
			expectedError:      false,
		},
		{
			name:                 "Empty Title",
			userID:               "user123",
			title:                "",
			description:          "Valid Description",
			category:             1,
			consultationFormat:   1,
			consultationMethod:   1,
			budget:               1000,
			period:               30,
			status:               1,
			expectedError:        true,
			expectedErrorMessage: "title must not be empty",
		},
		{
			name:                 "Empty Description",
			userID:               "user123",
			title:                "Valid Title",
			description:          "",
			category:             1,
			consultationFormat:   1,
			consultationMethod:   1,
			budget:               1000,
			period:               30,
			status:               1,
			expectedError:        true,
			expectedErrorMessage: "description must not be empty",
		},
		{
			name:                 "Title Too Long",
			userID:               "user123",
			title:                string(make([]byte, 256)), // 256文字のタイトル
			description:          "Valid Description",
			category:             1,
			consultationFormat:   1,
			consultationMethod:   1,
			budget:               1000,
			period:               30,
			status:               1,
			expectedError:        true,
			expectedErrorMessage: "title must be less than 255 characters",
		},
		{
			name:                 "Description Too Long",
			userID:               "user123",
			title:                "Valid Title",
			description:          string(make([]byte, 2001)), // 2001文字の説明
			category:             1,
			consultationFormat:   1,
			consultationMethod:   1,
			budget:               1000,
			period:               30,
			status:               1,
			expectedError:        true,
			expectedErrorMessage: "description must be less than 2000 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mentorRecruitment, err := mentordm.GenWhenCreate(
				tt.userID,
				tt.title,
				tt.description,
				tt.category,
				tt.consultationFormat,
				tt.consultationMethod,
				tt.budget,
				tt.period,
				tt.status,
			)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, mentorRecruitment)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, mentorRecruitment)
				assert.Equal(t, tt.userID, mentorRecruitment.UserID)
				assert.Equal(t, tt.title, mentorRecruitment.Title)
				assert.Equal(t, tt.description, mentorRecruitment.Description)
				assert.Equal(t, tt.category, mentorRecruitment.Category)
				assert.Equal(t, tt.consultationFormat, mentorRecruitment.ConsultationFormat)
				assert.Equal(t, tt.consultationMethod, mentorRecruitment.ConsultationMethod)
				assert.Equal(t, tt.budget, mentorRecruitment.Budget)
				assert.Equal(t, tt.period, mentorRecruitment.Period)
				assert.Equal(t, tt.status, mentorRecruitment.Status)
			}
		})
	}
}
