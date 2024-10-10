package contractoutput

type CreateContractOutput struct {
	ID                 string `json:"id,omitempty"`
	PlanID             string `json:"plan_id,omitempty"`
	UserID             string `json:"user_id,omitempty"`
	ContractApprovalID string `json:"contract_approval_id,omitempty"`
	Message            string `json:"message,omitempty"`
	Status             uint8  `json:"status,omitempty"`
}
