package contractrequestinput

type CreateContractRequestInput struct {
	ID      string `json:"id,omitempty"`
	PlanID  string `json:"plan_id,omitempty"`
	Message string `json:"message,omitempty"`
}
