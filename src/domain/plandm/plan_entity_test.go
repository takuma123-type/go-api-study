package plandm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlan(t *testing.T) {
	tests := []struct {
		name                 string
		userID               string
		title                string
		content              string
		category             int
		status               int
		consultationFormat   int
		price                int
		consultationMethod   int
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:               "Valid Plan",
			userID:             "user123",
			title:              "Valid Title",
			content:            "Valid Content",
			category:           1,
			status:             1,
			consultationFormat: 1,
			price:              100,
			consultationMethod: 1,
			expectedError:      false,
		},
		{
			name:                 "Empty Title",
			userID:               "user123",
			title:                "",
			content:              "Valid Content",
			category:             1,
			status:               1,
			consultationFormat:   1,
			price:                100,
			consultationMethod:   1,
			expectedError:        true,
			expectedErrorMessage: "title must not be empty",
		},
		{
			name:                 "Empty Content",
			userID:               "user123",
			title:                "Valid Title",
			content:              "",
			category:             1,
			status:               1,
			consultationFormat:   1,
			price:                100,
			consultationMethod:   1,
			expectedError:        true,
			expectedErrorMessage: "content must not be empty",
		},
		{
			name:                 "Title Too Long",
			userID:               "user123",
			title:                string(make([]rune, titleLength+1)),
			content:              "Valid Content",
			category:             1,
			status:               1,
			consultationFormat:   1,
			price:                100,
			consultationMethod:   1,
			expectedError:        true,
			expectedErrorMessage: "title must be less than 255 characters",
		},
		{
			name:                 "Content Too Long",
			userID:               "user123",
			title:                "Valid Title",
			content:              string(make([]rune, contentLength+1)),
			category:             1,
			status:               1,
			consultationFormat:   1,
			price:                100,
			consultationMethod:   1,
			expectedError:        true,
			expectedErrorMessage: "content must be less than 2000 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan, err := newPlan(tt.userID, tt.title, tt.content, tt.category, tt.status, tt.consultationFormat, tt.price, tt.consultationMethod)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, plan)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, plan)
				assert.Equal(t, tt.userID, plan.GetUserID())
				assert.Equal(t, tt.title, plan.GetTitle())
				assert.Equal(t, tt.content, plan.GetContent())
				assert.Equal(t, tt.category, plan.GetCategory())
				assert.Equal(t, tt.status, plan.GetStatus())
				assert.Equal(t, tt.consultationFormat, plan.GetConsultationFormat())
				assert.Equal(t, tt.price, plan.GetPrice())
				assert.Equal(t, tt.consultationMethod, plan.GetConsultationMethod())
			}
		})
	}
}
