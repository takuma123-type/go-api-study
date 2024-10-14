package contractapprovalinput

type CreateContractApprovalInput struct {
	ID      string `json:"id,omitempty"`
	PlanID  string `json:"plan_id,omitempty"`
	Message string `json:"message,omitempty"`
}
