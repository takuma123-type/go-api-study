package plandm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
)

func TestGenPlanIfCreate(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan, err := plandm.GenPlanIfCreate(tt.userID, tt.title, tt.content, tt.category, tt.consultationFormat, tt.consultationMethod, tt.price, tt.status)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, plan)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, plan)

				assert.Equal(t, tt.userID, plan.UserID())
				assert.Equal(t, tt.title, plan.Title())
				assert.Equal(t, tt.content, plan.Content())
				assert.Equal(t, uint16(tt.category), plan.Category())
				assert.Equal(t, uint16(tt.status), plan.Status())
				assert.Equal(t, uint16(tt.consultationFormat), plan.ConsultationFormat())
				assert.Equal(t, uint16(tt.price), plan.Price())
				assert.Equal(t, tt.consultationMethod, plan.ConsultationMethod())
			}
		})
	}
}
