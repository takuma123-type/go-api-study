package shared

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreatedAt(t *testing.T) {
	// time.Now() の結果を変数に格納
	now := time.Now()

	tests := []struct {
		name     string
		time1    time.Time
		time2    time.Time
		expected bool
	}{
		{
			name:     "Equal times should be equal",
			time1:    now,
			time2:    now,
			expected: true,
		},
		{
			name:     "Different times should not be equal",
			time1:    now,
			time2:    now.Add(1 * time.Second),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdAt1 := CreatedAt(tt.time1)
			createdAt2 := CreatedAt(tt.time2)
			assert.Equal(t, tt.expected, createdAt1.Equal(createdAt2))
		})
	}

	t.Run("NewCreatedAt should create a valid CreatedAt object", func(t *testing.T) {
		createdAt := NewCreatedAt()
		assert.WithinDuration(t, time.Now(), createdAt.Value(), time.Second, "Expected CreatedAt to be close to current time")
	})
}
