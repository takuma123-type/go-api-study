package contractdm

import "context"

type ContractRepository interface {
	Store(ctx context.Context, contract *Contract) error
}
