package contractdm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/contractdm"
)

func TestGenContractIfCreate(t *testing.T) {
	tests := []struct {
		name                 string
		id                   string
		planID               string
		userID               string
		contractApprovalID   string
		message              string
		status               uint8
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:               "Valid Contract",
			id:                 "contract123",
			planID:             "plan123",
			userID:             "user123",
			contractApprovalID: "approval123",
			message:            "Valid Contract Message",
			status:             1,
			expectedError:      false,
		},
		{
			name:                 "Empty PlanID",
			id:                   "contract124",
			planID:               "",
			userID:               "user123",
			contractApprovalID:   "approval123",
			message:              "Valid Contract Message",
			status:               1,
			expectedError:        true,
			expectedErrorMessage: "planID must not be empty",
		},
		{
			name:                 "Empty UserID",
			id:                   "contract125",
			planID:               "plan123",
			userID:               "",
			contractApprovalID:   "approval123",
			message:              "Valid Contract Message",
			status:               1,
			expectedError:        true,
			expectedErrorMessage: "userID must not be empty",
		},
		{
			name:                 "Empty ContractApprovalID",
			id:                   "contract126",
			planID:               "plan123",
			userID:               "user123",
			contractApprovalID:   "",
			message:              "Valid Contract Message",
			status:               1,
			expectedError:        true,
			expectedErrorMessage: "contractApprovalID must not be empty",
		},
		{
			name:                 "Empty Message",
			id:                   "contract127",
			planID:               "plan123",
			userID:               "user123",
			contractApprovalID:   "approval123",
			message:              "",
			status:               2,
			expectedError:        true,
			expectedErrorMessage: "message must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contract, err := contractdm.GenContractIfCreate(
				tt.id,
				tt.planID,
				tt.userID,
				tt.contractApprovalID,
				tt.message,
				tt.status,
			)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, contract)
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, contract)

				assert.Equal(t, tt.id, contract.ID().String())
				assert.Equal(t, tt.planID, contract.PlanID())
				assert.Equal(t, tt.userID, contract.UserID())
				assert.Equal(t, tt.contractApprovalID, contract.ContractApprovalID())
				assert.Equal(t, tt.message, contract.Message())
				assert.Equal(t, tt.status, contract.Status())
			}
		})
	}
}
