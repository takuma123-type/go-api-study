package contractapprovaloutput

type CreateContractApprovalOutput struct {
	ID                string `json:"id,omitempty"`
	ContractRequestID string `json:"contract_request_id,omitempty"`
	Message           string `json:"message,omitempty"`
}
