package plandm

import "context"

type PlanRepository interface {
	Store(ctx context.Context, plan *Plan) error
	UpdateStatus(ctx context.Context, planID string, status uint8) error
}
