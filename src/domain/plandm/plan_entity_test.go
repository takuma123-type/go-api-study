package plandm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/plandm"
)

func TestGenPlanIfCreate(t *testing.T) {
	tests := []struct {
		name                 string
		id                   string
		userID               string
		title                string
		content              string
		category             uint16
		status               uint16
		consultationFormat   uint16
		price                uint16
		consultationMethod   uint8
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:               "Valid Plan",
			id:                 "plan123",
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
			id:                   "plan124",
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
			id:                   "plan125",
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
			plan, err := plandm.GenPlanIfCreate(
				tt.id,
				tt.userID,
				tt.title,
				tt.content,
				int(tt.category),
				int(tt.status),
				int(tt.consultationFormat),
				int(tt.price),
				int(tt.consultationMethod),
			)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, plan)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, plan)

				assert.Equal(t, tt.id, plan.ID().String())
				assert.Equal(t, tt.userID, plan.UserID())
				assert.Equal(t, tt.title, plan.Title())
				assert.Equal(t, tt.content, plan.Content())
				assert.Equal(t, tt.category, plan.Category())
				assert.Equal(t, tt.status, plan.Status())
				assert.Equal(t, tt.consultationFormat, plan.ConsultationFormat())
				assert.Equal(t, tt.price, plan.Price())
				assert.Equal(t, tt.consultationMethod, plan.ConsultationMethod())
			}
		})
	}
}
