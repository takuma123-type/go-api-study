package contractapprovaldm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/contractapprovaldm"
)

func TestGenContractApprovalIfCreate(t *testing.T) {
	tests := []struct {
		name                 string
		contractRequestID    string
		message              string
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:              "Valid Contract Approval",
			contractRequestID: "contractReq123",
			message:           "This is a valid message",
			expectedError:     false,
		},
		{
			name:                 "Empty Contract Request ID",
			contractRequestID:    "",
			message:              "This is a valid message",
			expectedError:        true,
			expectedErrorMessage: "contractRequestIDは必須です",
		},
		{
			name:                 "Empty Message",
			contractRequestID:    "contractReq123",
			message:              "",
			expectedError:        true,
			expectedErrorMessage: "メッセージは必須です",
		},
		{
			name:                 "Message Too Long",
			contractRequestID:    "contractReq123",
			message:              "This message is way too long" + string(make([]byte, 501)),
			expectedError:        true,
			expectedErrorMessage: "メッセージは500文字以内で入力してください",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contractApproval, err := contractapprovaldm.GenContractApprovalIfCreate(
				tt.contractRequestID,
				tt.message,
			)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, contractApproval)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, contractApproval)

				assert.Equal(t, tt.contractRequestID, contractApproval.ContractRequestID().String())
				assert.Equal(t, tt.message, contractApproval.Message())
			}
		})
	}
}
