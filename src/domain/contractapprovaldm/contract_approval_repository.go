package contractapprovaldm

import "context"

type ContractApprovalRepository interface {
	Store(ctx context.Context, contractApproval *ContractApproval) error
}
