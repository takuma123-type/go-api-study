package shared

import "time"

type UpdatedAt time.Time

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func (u UpdatedAt) Value() time.Time {
	return time.Time(u)
}

func (u UpdatedAt) Equal(u2 UpdatedAt) bool {
	return u.Value().Equal(u2.Value())
}
