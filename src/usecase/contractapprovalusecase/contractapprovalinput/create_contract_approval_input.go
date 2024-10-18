package contractapprovalinput

type CreateContractApprovalInput struct {
	ID                string `json:"id,omitempty"`
	ContractRequestID string `json:"contract_request_id,omitempty"`
	Message           string `json:"message,omitempty"`
}
