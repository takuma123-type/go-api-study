package plandm

import "context"

type PlanRepository interface {
	Store(ctx context.Context, plan *Plan) error
}
