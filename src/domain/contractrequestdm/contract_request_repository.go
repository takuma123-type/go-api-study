package contractrequestdm

import "context"

type ContractRequestRepository interface {
	Store(ctx context.Context, contractRequest *ContractRequest) error
}
