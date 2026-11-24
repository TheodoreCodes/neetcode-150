package slidingWindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem constraints:
//
//	1 <= prices.length <= 100
//	0 <= prices[i] <= 100
func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"example 1", []int{10, 1, 5, 6, 7, 1}, 6},
		{"example 2", []int{10, 8, 7, 5, 2}, 0},
		{"single day", []int{5}, 0},
		{"two days rising", []int{1, 2}, 1},
		{"two days falling", []int{2, 1}, 0},
		{"all equal", []int{3, 3, 3}, 0},
		{"strictly increasing", []int{1, 2, 3, 4, 5}, 4},
		{"global min is last", []int{4, 6, 8, 0}, 4},
		{"best buy is not global min", []int{5, 100, 0, 3}, 95},
		{"best sell before a later dip", []int{2, 9, 1, 4}, 7},
		{"later pair beats earlier pair", []int{1, 3, 0, 8}, 8},
		{"zero price", []int{0, 100}, 100},
		{"extreme values", []int{100, 0, 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProfit(tt.prices))
		})
	}
}

func TestMaxProfitMaxLength(t *testing.T) {
	prices := make([]int, 100)
	for i := range prices {
		prices[i] = i
	}
	assert.Equal(t, 99, maxProfit(prices))
}
