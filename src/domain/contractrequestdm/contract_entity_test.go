package contractrequestdm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/contractrequestdm"
)

func TestGenContractRequestIfCreate(t *testing.T) {
	tests := []struct {
		name                 string
		planID               string
		message              string
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:          "Valid Contract Request",
			planID:        "plan123",
			message:       "This is a valid message",
			expectedError: false,
		},
		{
			name:                 "Empty Plan ID",
			planID:               "",
			message:              "This is a valid message",
			expectedError:        true,
			expectedErrorMessage: "planIDは必須です",
		},
		{
			name:                 "Empty Message",
			planID:               "plan123",
			message:              "",
			expectedError:        true,
			expectedErrorMessage: "メッセージは必須です",
		},
		{
			name:                 "Message Too Long",
			planID:               "plan123",
			message:              "This message is way too long" + string(make([]byte, 501)), // 501文字のメッセージ
			expectedError:        true,
			expectedErrorMessage: "メッセージは500文字以内で入力してください",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contractRequest, err := contractrequestdm.GenContractRequestIfCreate(
				tt.planID,
				tt.message,
			)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, contractRequest)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, contractRequest)

				assert.Equal(t, tt.planID, contractRequest.PlanID().String())
				assert.Equal(t, tt.message, contractRequest.Message())
			}
		})
	}
}
