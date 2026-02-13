package taxi_time

import (
	"testing"
	"time"
)

func TestGetTimeMultiplier(t *testing.T) {
	cases := []struct {
		name string
		value time.Time
		expected float64
	} {
		{
			name: "Ночной тариф",
			value: time.Date(2026, 2, 12, 3, 0, 0, 0, time.UTC),
			expected: 1.5,
		},
		{
			name: "Утренний час пик",
			value: time.Date(2026, 2, 12, 8, 0, 0, 0, time.UTC),
			expected: 1.3,
		},
		{
			name: "Выходные",
			value: time.Date(2026, 2, 14, 15, 0, 0, 0, time.UTC),
			expected: 1.2,
		},
		{
			name: "Обычный тариф",
			value: time.Date(2026, 2, 12, 15, 0, 0, 0, time.UTC),
			expected: 1.,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GetTimeMultiplier(c.value)
			if got != c.expected {
				t.Errorf("GetTimeMultiplier(%v) = %v; expected %v", c.value, got, c.expected)
			}
		})
	}
}