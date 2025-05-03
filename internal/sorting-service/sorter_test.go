package sortingservice

import (
	"testing"
)

func TestSortAndCompareSlices(t *testing.T) {
	tests := []struct {
		name           string
		containers     []int
		balls          []int
		expectedResult string
	}{
		{
			name:           "Equal sorted slices",
			containers:     []int{1, 2, 3},
			balls:          []int{1, 2, 3},
			expectedResult: "yes",
		},
		{
			name:           "Equal unsorted slices",
			containers:     []int{3, 2, 1},
			balls:          []int{1, 2, 3},
			expectedResult: "yes",
		},
		{
			name:           "Balls do not fit into containers",
			containers:     []int{1, 2, 3},
			balls:          []int{1, 2, 20},
			expectedResult: "no",
		},
		{
			name:           "Same sums but different distributions",
			containers:     []int{1, 2, 3},
			balls:          []int{2, 2, 2},
			expectedResult: "no",
		},
		{
			name:           "Single container and same color balls match",
			containers:     []int{5},
			balls:          []int{5},
			expectedResult: "yes",
		},
		{
			name:           "Single container and same color balls do not match",
			containers:     []int{5},
			balls:          []int{10},
			expectedResult: "no",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortAndCompareSlices(tt.containers, tt.balls)
			if result != tt.expectedResult {
				t.Errorf("SortAndCompareSlices() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}
